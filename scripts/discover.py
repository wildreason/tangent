#!/usr/bin/env python3
"""
Rapid Design Discovery System
Usage: python discover.py --agent ga --state search --iterations 10
"""

import json
import subprocess
import shutil
import random
import copy
from pathlib import Path
from datetime import datetime

class DesignDiscovery:
    def __init__(self, agent, state, state_path, project_root):
        self.agent = agent
        self.state = state
        self.state_path = Path(state_path)
        self.project_root = Path(project_root)
        self.backup_path = self.state_path.with_suffix('.backup')
        self.approved = []
        self.session_id = datetime.now().strftime('%Y%m%d_%H%M%S')
        
    def backup(self):
        """Create backup of original state"""
        shutil.copy(self.state_path, self.backup_path)
        print(f"💾 Backup created: {self.backup_path}")
    
    def restore(self):
        """Restore from backup"""
        if self.backup_path.exists():
            shutil.copy(self.backup_path, self.state_path)
    
    def mutate_design(self, data):
        """Generate random design variation"""
        mutations = [
            self._remove_random_frames,
            self._modify_line,
            self._duplicate_frame,
            self._shuffle_frames,
            self._modify_multiple_lines
        ]
        
        # Pick 1-2 mutations for more variety
        num_mutations = random.randint(1, 2)
        selected = random.sample(mutations, num_mutations)
        
        print(f"   Applying {num_mutations} mutation(s):")
        for mutation in selected:
            print(f"   - {mutation.__name__.replace('_', ' ')}")
            data = mutation(data)
        
        return data
    
    def _remove_random_frames(self, data):
        """Remove random frames, keep at least 2"""
        if len(data['frames']) > 3:
            keep = random.randint(2, len(data['frames']) - 1)
            original_count = len(data['frames'])
            data['frames'] = random.sample(data['frames'], keep)
            print(f"      Frames: {original_count} → {keep}")
        return data
    
    def _modify_line(self, data):
        """Modify a random line in a random frame"""
        if data['frames']:
            frame_idx = random.randint(0, len(data['frames']) - 1)
            frame = data['frames'][frame_idx]
            if frame['lines']:
                line_idx = random.randint(0, len(frame['lines']) - 1)
                old_line = frame['lines'][line_idx]
                frame['lines'][line_idx] = self._random_pattern()
                print(f"      Line {line_idx} in frame {frame_idx}: modified")
        return data
    
    def _modify_multiple_lines(self, data):
        """Modify 2-3 lines across frames"""
        if data['frames']:
            num_mods = random.randint(2, 3)
            for _ in range(num_mods):
                frame = random.choice(data['frames'])
                if frame['lines']:
                    idx = random.randint(0, len(frame['lines']) - 1)
                    frame['lines'][idx] = self._random_pattern()
            print(f"      Modified {num_mods} lines")
        return data
    
    def _random_pattern(self):
        """Generate random pattern string"""
        patterns = [
            '_rfffffffl_',
            '_fffffffff_',
            '___________',
            '_r_______l_',
            '_rrrrrrrrrl_',
            '_rf______l_',
            '_r__fff__l_',
            '_rff___ffl_',
            '           ',  # Empty
            '_r_fff___l_'
        ]
        return random.choice(patterns)
    
    def _duplicate_frame(self, data):
        """Duplicate a random frame"""
        if data['frames']:
            frame_to_dup = random.choice(data['frames'])
            data['frames'].append(copy.deepcopy(frame_to_dup))
            print(f"      Duplicated frame (total: {len(data['frames'])})")
        return data
    
    def _shuffle_frames(self, data):
        """Shuffle frame order"""
        original_count = len(data['frames'])
        random.shuffle(data['frames'])
        print(f"      Shuffled {original_count} frames")
        return data
    
    def build(self):
        """Build the CLI using make"""
        result = subprocess.run(
            ['make', 'build-cli'],
            cwd=self.project_root,
            capture_output=True,
            text=True
        )
        return result.returncode == 0, result.stderr
    
    def preview(self):
        """Show animation to user"""
        cli_path = self.project_root / 'tangent-cli'
        subprocess.run([
            str(cli_path), 'browse', self.agent,
            '--state', self.state,
            '--fps', '3',
            '--loops', '2'
        ], cwd=self.project_root)
    
    def prompt_approval(self, demo_mode=False):
        """Ask user for approval"""
        if demo_mode:
            # In demo mode, randomly approve or decline
            choice = random.choice(['a', 'a', 'd'])  # 66% approve, 33% decline
            print(f"\n[DEMO MODE] Automatically choosing: {choice.upper()}")
            return choice
        
        while True:
            choice = input("\n[A]pprove  [D]ecline  [Q]uit: ").strip().lower()
            if choice in ['a', 'd', 'q']:
                return choice
            print("❌ Invalid choice. Use A, D, or Q")
    
    def save_approved(self):
        """Save all approved designs to file"""
        if not self.approved:
            return
        
        output_dir = self.project_root / 'testing' / 'discovery'
        output_dir.mkdir(parents=True, exist_ok=True)
        
        output_file = output_dir / f'{self.state}_{self.session_id}_approved.json'
        with open(output_file, 'w') as f:
            json.dump(self.approved, f, indent=2)
        
        print(f"💾 Saved approved designs to: {output_file}")
    
    def run(self, iterations, demo_mode=False):
        """Main discovery loop"""
        print(f"╔{'═'*60}╗")
        print(f"║  🔍 DESIGN DISCOVERY SYSTEM                                ║")
        print(f"╚{'═'*60}╝")
        print(f"Agent: {self.agent}")
        print(f"State: {self.state}")
        print(f"Iterations: {iterations}")
        if demo_mode:
            print(f"Mode: DEMO (automatic decisions)")
        print()
        
        self.backup()
        
        try:
            for i in range(iterations):
                print(f"\n{'─'*60}")
                print(f"🎨 ITERATION {i+1}/{iterations}")
                print(f"{'─'*60}")
                
                # Load current state
                with open(self.state_path) as f:
                    data = json.load(f)
                
                original_data = copy.deepcopy(data)
                
                # Mutate
                print("🔧 Generating variation...")
                mutated = self.mutate_design(data)
                
                # Save temporary version
                with open(self.state_path, 'w') as f:
                    json.dump(mutated, f, indent=2)
                
                # Build
                print("\n🔨 Building...")
                success, error = self.build()
                if not success:
                    print(f"❌ Build failed:")
                    print(f"   {error[:200]}")
                    print("   Reverting and continuing...")
                    self.restore()
                    continue
                print("✅ Build successful")
                
                # Preview
                print("\n📺 Preview:")
                print(f"{'─'*60}")
                self.preview()
                print(f"{'─'*60}")
                
                # Approve
                choice = self.prompt_approval(demo_mode)
                
                if choice == 'q':
                    print("\n👋 Quitting discovery session...")
                    self.restore()
                    break
                elif choice == 'a':
                    print("✅ APPROVED! Keeping this design as base for next iteration")
                    self.approved.append({
                        'iteration': i+1,
                        'timestamp': datetime.now().isoformat(),
                        'mutations': len(mutated.get('frames', [])),
                        'data': mutated
                    })
                    # Keep mutated version as new base
                else:
                    print("❌ DECLINED - Reverting to previous design")
                    with open(self.state_path, 'w') as f:
                        json.dump(original_data, f, indent=2)
            
            # Summary
            print(f"\n{'═'*60}")
            print(f"📊 DISCOVERY SESSION COMPLETE")
            print(f"{'═'*60}")
            print(f"Total iterations: {iterations}")
            print(f"✅ Approved designs: {len(self.approved)}")
            print(f"❌ Declined designs: {i + 1 - len(self.approved)}")
            
            if self.approved:
                print(f"\n💡 Your current state contains the last approved design.")
                print(f"📁 Original backup: {self.backup_path}")
                self.save_approved()
            else:
                print(f"\n⚠️  No designs approved. Restoring original...")
                self.restore()
                
        except KeyboardInterrupt:
            print("\n\n⚠️  Interrupted by user. Restoring original state...")
            self.restore()
            print("✅ Original state restored.")
        except Exception as e:
            print(f"\n❌ Error occurred: {e}")
            print("🔄 Restoring original state...")
            self.restore()
            raise

def main():
    import argparse
    parser = argparse.ArgumentParser(
        description='Design Discovery System - Explore animation designs through rapid iteration'
    )
    parser.add_argument('--agent', default='ga', help='Agent character to use (default: ga)')
    parser.add_argument('--state', default='search', help='State to explore (default: search)')
    parser.add_argument('--iterations', type=int, default=10, help='Number of iterations (default: 10)')
    parser.add_argument('--demo', action='store_true', help='Demo mode: automatically approve/decline (for testing)')
    args = parser.parse_args()
    
    # Determine project root (script is in scripts/ directory)
    script_dir = Path(__file__).parent
    project_root = script_dir.parent
    
    state_path = project_root / 'pkg' / 'characters' / 'stateregistry' / 'states' / f'{args.state}.json'
    
    if not state_path.exists():
        print(f"❌ Error: State file not found: {state_path}")
        print(f"\nAvailable states:")
        states_dir = state_path.parent
        for state_file in sorted(states_dir.glob('*.json')):
            if not state_file.name.endswith('-legacy.json'):
                print(f"  - {state_file.stem}")
        return 1
    
    discovery = DesignDiscovery(args.agent, args.state, state_path, project_root)
    discovery.run(args.iterations, demo_mode=args.demo)
    return 0

if __name__ == '__main__':
    exit(main())

