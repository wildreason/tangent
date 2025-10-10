package characters

import (
	"fmt"
	"strings"

	"github.com/wildreason/tangent/pkg/characters/domain"
)

// ErrorHandler provides centralized error handling with suggestions
type ErrorHandler struct{}

// NewErrorHandler creates a new error handler
func NewErrorHandler() *ErrorHandler {
	return &ErrorHandler{}
}

// HandleError processes an error and provides user-friendly output with suggestions
func (h *ErrorHandler) HandleError(context string, err error) {
	fmt.Printf("✗ %s: %v\n", context, err)

	// Get suggestion based on error type
	suggestion := domain.GetErrorSuggestion(err)
	if suggestion != "" {
		fmt.Printf("  ◢ Tip: %s\n", suggestion)
	}

	// Provide additional context-specific suggestions
	h.provideContextualSuggestions(context, err)
}

// provideContextualSuggestions provides additional suggestions based on context
func (h *ErrorHandler) provideContextualSuggestions(context string, err error) {
	contextLower := strings.ToLower(context)

	switch {
	case strings.Contains(contextLower, "create"):
		h.suggestCharacterCreation(err)
	case strings.Contains(contextLower, "load"):
		h.suggestCharacterLoading(err)
	case strings.Contains(contextLower, "save"):
		h.suggestCharacterSaving(err)
	case strings.Contains(contextLower, "animate"):
		h.suggestAnimation(err)
	case strings.Contains(contextLower, "pattern"):
		h.suggestPattern(err)
	default:
		h.suggestGeneral(err)
	}
}

// suggestCharacterCreation provides suggestions for character creation errors
func (h *ErrorHandler) suggestCharacterCreation(err error) {
	switch e := err.(type) {
	case *domain.ValidationError:
		switch e.Field {
		case "name":
			fmt.Println("  ◢ Try: characters.NewCharacterBuilderV2(\"my-robot\", 8, 4)")
		case "dimensions":
			fmt.Println("  ◢ Try: width: 8, height: 4 (common character sizes)")
		case "frames":
			fmt.Println("  ◢ Try: .AddFrame(\"idle\", []string{\"FRF\", \"LRL\", \"FRF\"})")
		}
	case *domain.PatternCompilationError:
		fmt.Println("  ◢ Try: Use valid pattern characters like F, R, L, T, B, 1-8, _")
	}
}

// suggestCharacterLoading provides suggestions for character loading errors
func (h *ErrorHandler) suggestCharacterLoading(err error) {
	switch e := err.(type) {
	case *domain.CharacterNotFoundError:
		fmt.Printf("  ◢ Try: characters.ListLibrary() to see available characters\n")
		fmt.Printf("  ◢ Or: Create a new character with name '%s'\n", e.Name)
	}
}

// suggestCharacterSaving provides suggestions for character saving errors
func (h *ErrorHandler) suggestCharacterSaving(err error) {
	fmt.Println("  ◢ Check: File permissions and disk space")
	fmt.Println("  ◢ Try: Ensure character name is valid and not empty")
}

// suggestAnimation provides suggestions for animation errors
func (h *ErrorHandler) suggestAnimation(err error) {
	switch e := err.(type) {
	case *domain.AnimationError:
		switch e.Operation {
		case "start":
			fmt.Println("  ◢ Try: Ensure character has frames and terminal supports ANSI")
		case "frame_display":
			fmt.Println("  ◢ Try: Check frame content and terminal width")
		case "timing":
			fmt.Println("  ◢ Try: Use fps between 1-30 and positive loop count")
		}
	}
}

// suggestPattern provides suggestions for pattern errors
func (h *ErrorHandler) suggestPattern(err error) {
	fmt.Println("  ◢ Valid characters: F=█, R=▐, L=▌, T=▀, B=▄, 1-8=quadrants, _=space")
	fmt.Println("  ◢ Example: \"FRF\" becomes \"█▐█\"")
}

// suggestGeneral provides general suggestions
func (h *ErrorHandler) suggestGeneral(err error) {
	fmt.Println("  ◢ Check: All required parameters are provided")
	fmt.Println("  ◢ Verify: Input values are within valid ranges")
	fmt.Println("  ◢ Review: Error message for specific details")
}

// FormatError formats an error with context and suggestions
func FormatError(context string, err error) string {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("✗ %s: %v\n", context, err))

	suggestion := domain.GetErrorSuggestion(err)
	if suggestion != "" {
		sb.WriteString(fmt.Sprintf("  ◢ Tip: %s\n", suggestion))
	}

	return sb.String()
}

// IsValidationError checks if an error is a validation error
func IsValidationError(err error) bool {
	_, ok := err.(*domain.ValidationError)
	return ok
}

// IsCharacterNotFoundError checks if an error is a character not found error
func IsCharacterNotFoundError(err error) bool {
	_, ok := err.(*domain.CharacterNotFoundError)
	return ok
}

// IsPatternCompilationError checks if an error is a pattern compilation error
func IsPatternCompilationError(err error) bool {
	_, ok := err.(*domain.PatternCompilationError)
	return ok
}

// IsAnimationError checks if an error is an animation error
func IsAnimationError(err error) bool {
	_, ok := err.(*domain.AnimationError)
	return ok
}

// GetErrorField returns the field name for validation errors
func GetErrorField(err error) string {
	if ve, ok := err.(*domain.ValidationError); ok {
		return ve.Field
	}
	return ""
}

// GetErrorValue returns the value that caused the error
func GetErrorValue(err error) interface{} {
	if ve, ok := err.(*domain.ValidationError); ok {
		return ve.Value
	}
	return nil
}
