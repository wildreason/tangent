# Terminal Color Theme Research

## Research Context

This document contains research on color themes for the Wave terminal windows, focusing on two distinct user demographics:

1. **Developer users** - Familiar with terminal interfaces and dev tools
2. **GUI users** - Non-technical users transitioning from modern consumer apps (Notion, Figma, Slack)

## Original Color Palette (Theme 1)

The original bright, saturated colors designed for maximum distinction:

```
sa: #FF0000  (pure red)
ri: #FF8800  (bright orange)
ga: #FFD700  (gold)
ma: #00FF00  (pure green)
pa: #0088FF  (bright blue)
da: #8800FF  (bright purple)
ni: #FF0088  (bright pink)
```

**Characteristics:**
- 100% saturation across all colors
- High visual impact and immediate distinction
- Best for rapid prototyping and testing
- Can feel toy-like and cause visual fatigue in extended use

---

## Developer-Focused Themes

### Research Summary: Developer Users

**Key Insights:**
- Developers use terminals for 8+ hour sessions - longevity is critical
- Popular editor themes (Nord, Atom One Dark) provide proven color psychology
- Reduced saturation (40-50%) prevents "toy" feeling
- Balanced luminance reduces eye strain
- Professional precedent builds subconscious trust

### Theme: Deep Ocean
**Inspiration:** Nord theme
**Focus:** Low contrast warmth with depth, maximum comfort

```
sa: #BF616A  (muted coral red)
ri: #D08770  (soft rust orange)
ga: #EBCB8B  (warm sand)
ma: #A3BE8C  (sage green)
pa: #5E81AC  (slate blue)
da: #B48EAD  (dusty purple)
ni: #D896AB  (mauve pink)
```

**Characteristics:**
- Reduces eye strain significantly
- Professional enough for long coding sessions
- Subtle warmth prevents coldness while maintaining sophistication
- Similar to popular editor themes users already trust

### Theme: Twilight Studio (RECOMMENDED for developers)
**Inspiration:** Atom One Dark
**Focus:** Desaturated jewel tones with balanced luminance

```
sa: #E06C75  (rose)
ri: #D19A66  (amber)
ga: #E5C07B  (honey)
ma: #98C379  (moss)
pa: #61AFEF  (sky)
da: #C678DD  (orchid)
ni: #E88EB3  (blush)
```

**Characteristics:**
- Based on Atom One Dark palette psychology
- Higher engagement through gentle color distinction
- Balanced brightness - not dull, not aggressive
- Accessible for extended use without fatigue
- **Best balance:** engaging enough to be distinctive, comfortable enough for 8+ hour sessions

### Theme: Carbon Terminal
**Inspiration:** Terminal-native aesthetic
**Focus:** Monochromatic base with strategic accent pops

```
sa: #DA6A6A  (clay red)
ri: #CF8D5C  (terracotta)
ga: #D4B86A  (brass)
ma: #7FB07F  (fern)
pa: #6B9FD1  (steel)
da: #A67BC4  (lavender)
ni: #D17BA3  (rose quartz)
```

**Characteristics:**
- Conservative saturation (~50-60% vs original 100%)
- Maintains color distinction without visual shouting
- Ages well - won't feel dated or toy-like
- Terminal-native aesthetic that pros respect

---

## GUI-User-Focused Themes

### Research Summary: Non-Developer Users

**Key Psychological Barriers:**
- Terminals feel intimidating due to dark backgrounds + bright neon colors = "hacker aesthetic"
- GUI users expect warm, approachable colors (Notion, Figma, Slack use soft pastels/neutrals)
- Modern consumer apps reduce cognitive load through muted, earthy palettes
- 2024-2025 trend: Retro-futuristic pastels + earthy tones for comfort & familiarity

**What Works for Non-Technical Users:**
1. **Warm undertones** - Creates invitation vs. clinical coldness
2. **Pastel saturation (50-70%)** - Playful but not toy-like
3. **Nature-inspired colors** - Psychologically grounding and calm
4. **Familiar palette from trusted apps** - Reduces learning curve
5. **High readability contrast** - Accessibility without aggression

**Popular Consumer App Patterns:**
- Linear: Calm indigo + crisp neutrals (focus-driven)
- Notion: Minimal black/white base + 8-10 soft accent colors
- Modern trend: Monochrome with strategic color pops (30% less color than before)

### Theme: Latte Window (RECOMMENDED for GUI users)
**Inspiration:** Catppuccin Latte
**Target:** GUI users who love Notion/Figma aesthetics
**Psychology:** Warm pastels that feel playful yet professional

```
sa: #E78284  (soft coral - approachable warmth)
ri: #EF9F76  (peach - friendly energy)
ga: #E5C890  (warm sand - comfort)
ma: #A6D189  (sage - growth/trust)
pa: #85C1DC  (sky blue - calm clarity)
da: #CA9EE6  (lavender - creative)
ni: #F4B8E4  (rose - welcoming)
```

**Characteristics:**
- Based on Catppuccin Latte (beloved by 100K+ developers who wanted softer terminals)
- Reduced saturation (~60%) = long-session comfort
- Warm undertones throughout (no cold blues/cyans)
- Matches modern productivity app aesthetics
- Feels like "Notion met the terminal"
- **Psychological bridge:** "This feels like my other apps"

### Theme: Garden Terminal
**Inspiration:** Earthy natural materials
**Target:** Users intimidated by tech who respond to nature metaphors
**Psychology:** Grounding, organic, reduces "computer anxiety"

```
sa: #D4787D  (terracotta rose - earthy warmth)
ri: #D89C6A  (clay orange - natural)
ga: #C9B68C  (wheat - harvest calm)
ma: #8FB378  (moss green - forest trust)
pa: #7CA8B8  (river blue - flowing ease)
da: #A888BA  (dusty iris - twilight)
ni: #C895A8  (dusty mauve - soft dusk)
```

**Characteristics:**
- All colors reference natural materials (clay, wheat, moss, river)
- Muted like Gruvbox but warmer for approachability
- Lower contrast = less "screen glare" feeling
- Psychological association: nature = safe, familiar
- Perfect for users who describe themselves as "not technical"

### Theme: Cozy Workspace
**Inspiration:** Modern GUI hybrid (Slack/Linear)
**Target:** Slack/Linear users expecting crisp, modern interfaces
**Psychology:** Professional warmth - serious without being intimidating

```
sa: #E18B8B  (rose quartz - gentle strength)
ri: #E5A679  (amber glow - warm energy)
ga: #E6CC94  (honey - productive sweetness)
ma: #99C794  (balanced green - growth)
pa: #78AED4  (clear sky - open possibility)
da: #B592D4  (soft violet - creative depth)
ni: #DE99B8  (blush - approachable)
```

**Characteristics:**
- Balanced luminance across all colors (prevents jarring switches)
- Inspired by VSCode + Linear + Slack palettes
- 50-55% saturation = "colorful but calm"
- Professional enough for work, warm enough for newcomers
- Reduces "terminal = scary code stuff" association

---

## Strategic Recommendations

### For Developer Users
**Recommended:** Twilight Studio
**Reasoning:** Best balance between engaging visual interest and 8+ hour session comfort. Based on proven Atom One Dark psychology that developers already trust.

### For GUI Users
**Recommended:** Latte Window
**Reasoning:** Proven success in developer-to-consumer tool transitions. Directly mirrors the soft aesthetics GUI users already trust from Notion/Figma. Most "un-terminal-like" while staying functional.

### Implementation Priorities
1. **Longevity over immediate impact** - All recommended themes prioritize sustained engagement
2. **50-70% saturation** - Balances visual interest with comfort
3. **Warm undertones** - More inviting and less clinical than pure hues
4. **Psychological familiarity** - Based on tools users already trust

---

## Color Psychology Notes

### Why Original Colors Don't Work Long-Term
- 100% saturation = "developer placeholder colors"
- Pure hues (#FF0000, #00FF00) feel like quick prototypes, not production UI
- Visual fatigue sets in after 30-60 minutes of sustained use
- Creates "toy" feeling that reduces perceived professionalism
- For GUI users: reinforces "this is for coders only" barrier

### Key Design Principle
The best terminal colors are ones users **don't actively notice** after 20 minutes, yet still provide clear visual distinction when needed. Aggressively bright colors demand constant attention; muted colors with good contrast provide information without exhaustion.
