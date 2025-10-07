# Paris Character Example

This example demonstrates a character created using the `tangent` CLI builder and exported as Go code.

## Character Details

- **Name**: paris
- **Dimensions**: 11x3
- **Frames**: 3 (head, wave, wave-2)
- **Animation**: Waving motion

## How It Was Created

1. Run `tangent` to launch the interactive builder
2. Create a new character named "paris" with dimensions 11x3
3. Add frames using pattern codes:
   - **head**: Idle position
   - **wave**: Right arm extended
   - **wave-2**: Left arm extended
4. Export the code using option 5
5. Copy the exported code into this example

## Pattern Codes Used

```
Frame: head
__R6FFF6L__
_T6FFFFF5T_
___11_22___

Frame: wave
__R6FFF6L__
__6FFFFF5T1
___11_22___

Frame: wave-2
__R6FFF6L__
2T6FFFFF5__
___11_22___
```

## Running

```bash
cd examples/tokyo
go run main.go
```

## What It Does

1. Builds the character from the pattern spec
2. Shows the idle frame
3. Animates through all frames at 5 FPS for 5 cycles

---
*Created with Tangent v0.0.1 - Terminal Agent Designer*  
*© 2025 Wildreason, Inc*

