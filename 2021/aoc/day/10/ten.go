package day

import "sort"

type stack []rune

func Ten(input []string) (int, int) {
	sum := 0

	// needed for part 2
	var incompleteStacks []stack

	for _, row := range input {
		// a placeholder char so we can see if our row is corrupted
		foundInstead := '?'
		s := stack{}

		for _, char := range row {
			if foundInstead != '?' { break }

			lc := s.peek()

			switch char {
			case '(','[', '{', '<':
				s.push(char)
			case ')':
				if lc == '(' {
					s = s.pop()
				} else {
					foundInstead = char
				}
			case ']':
				if lc == '[' {
					s = s.pop()
				} else {
					foundInstead = char
				}
			case '}':
				if lc == '{' {
					s = s.pop()
				} else {
					foundInstead = char
				}
			case '>':
				if lc == '<' {
					s = s.pop()
				} else {
					foundInstead = char
				}
			}

			switch foundInstead {
			case ')':
				sum += 3
			case ']':
				sum += 57
			case '}':
				sum += 1197
			case '>':
				sum += 25137
			}
		}

		// if the row isn't corrupted and has unmatched runes
		if foundInstead == '?' && len(s) > 0 {
			incompleteStacks = append(incompleteStacks, s)
		}
	}

	var scores []int

	for _, s := range incompleteStacks {
		score := 0
		// traverse the stack of unmatched runes from right to left
		for i := len(s) - 1; i >= 0; i-- {
			score *= 5
			switch s[i] {
			case '(':
				score += 1
			case '[':
				score += 2
			case '{':
				score += 3
			case '<':
				score += 4
			}
		}

		scores = append(scores, score)
	}

	sort.Ints(scores)

	return sum, scores[len(scores)/2]
}

func (s *stack) push(char rune) {
	*s = append(*s, char)
}

// couldn't figure out a way to make this work with
// a pointer to the original stack :(
func (s stack) pop() stack {
	return s[:len(s)-1]
}

func (s stack) peek() rune {
	if len(s) > 0 {
		return s[len(s)-1]
	}
	// a placeholder char as I couldn't be bothered to
	// mess around with nil pointer references
	return '0'
}
