package leetcode

import (
    "math"
    "unicode/utf8"
)

// Link: https://leetcode.com/problems/string-to-integer-atoi/

// Implement the myAtoi(string s) function, which converts a string to a 32-bit signed integer 
// (similar to C/C++'s atoi function).

// The algorithm for myAtoi(string s) is as follows:
//   1. Read in and ignore any leading whitespace.
//   2. Check if the next character (if not already at the end of the string) is '-' or '+'. Read 
//      this character in if it is either. This determines if the final result is negative or 
//      positive respectively. Assume the result is positive if neither is present.
//   3. Read in next the characters until the next non-digit charcter or the end of the input is 
//      reached. The rest of the string is ignored.
//   4. Convert these digits into an integer (i.e. "123" -> 123, "0032" -> 32). If no digits were
//      read, then the integer is 0. Change the sign as necessary (from step 2).
//   5. If the integer is out of the 32-bit signed integer range [-231, 231 - 1], then clamp the 
//      integer so that it remains in the range. Specifically, integers less than -231 should be 
//      clamped to -231, and integers greater than 231 - 1 should be clamped to 231 - 1.
//   6. Return the integer as the final result.

// Note:
//  - Only the space character ' ' is considered a whitespace character.
//  - Do not ignore any characters other than the leading whitespace or the rest of the string 
//    after the digits.

// Constraints:
//   - 0 <= s.length <= 200
//   - s consists of English letters (lower-case and upper-case), digits (0-9), ' ', '+', '-', 
//     and '.'.
func myAtoi(s string) int {
    digits := newAtoiDigits(len(s))

    b := []byte(s)
    for len(b) > 0 {
		char, size := getCharAndSize(b)
        nextChar := getNextChar(b, size)

        // Shortcut: 1e11 and beyond is already beyond the 32 bit 
        // bounds regardless of the signage
        if digits.len() > 10 {
            break
        } else if isWhiteSpace(char) {
            if !isWhiteSpace(nextChar) && !isDigit(nextChar) && !isSign(nextChar) {
                return 0
            }
        } else if isSign(char) {
            if isDigit(nextChar) {
                digits.setSign(toSignScalar(char))
            } else {
                return 0
            }
        } else if isDigit(char) {
            digits.push(toDigit(char))
            if !isDigit(nextChar) {
                break
            }
        } else {
            return 0
        }

        b = b[size:]
	}

    return digits.sum()
}

const maxInt32 = (1 << 31) - 1
const minInt32 = -maxInt32 - 1

// getCharAndSize returns the first encoded character and
// the number of bytes used in the encoding for that character.
func getCharAndSize(b []byte) (string, int) {
    r, size := utf8.DecodeRune(b)
    return string(r), size
}

// getNextChar gets the second encoded character based off the
// given number of bytes that the first character was encoded with.
func getNextChar(b []byte, size int) string {
    r, _ := utf8.DecodeRune(b[size:])
    if r == utf8.RuneError {
        return ""
    }
    return string(r)
}

// newAtoiDigits returns a pointer to a new instance of atoiDigits
func newAtoiDigits(capacity int) *atoiDigits {
    return &atoiDigits{values: make([]uint, 0, capacity), signScalar: 1}
}

// atoiDigits is a structure for tracking digits used in the "String to Integer Atoi"
// LeetCode problem.
type atoiDigits struct {
    values     []uint
    signScalar int8
}

// push appends the given number into the tracked digits,
// but only if it is not a leading zero.
func (d *atoiDigits) push(num uint) {
    if len(d.values) != 0 || num != 0 {
        d.values = append(d.values, num)
    }
}

// sum returns the decimal expansion sum of the tracked digits.
// This is bounded by the upper and lower limits of 32 bit
// integers.
func (d *atoiDigits) sum() int {
    sum := 0
    for i, val := range d.values {
        if val == 0 {
            continue
        }
        exp := d.len() - (i+1)
        if exp > 10 {
            sum = maxInt32+1
            break
        }
        inc := float64(val) * math.Pow(10, float64(exp))
        sum += int(inc)
    }
    number := sum * int(d.sign())
    if number < minInt32 {
		return minInt32
	}
    if number > maxInt32 {
		return maxInt32
	}
    return number
}

// setSign assigns the given sign scalar to this digits struct.
func (d *atoiDigits) setSign(scalar int8) {
    if scalar != 1 && scalar != -1 {
        panic("invalid scalar")
    }
    d.signScalar = scalar
}

// sign returns the tracked sign scalar.
func (d *atoiDigits) sign() int8 {
    return d.signScalar
}

// len returns the number of digits tracked by this digits struct.
func (d *atoiDigits) len() int {
    return len(d.values)
}

// toDigit converts a character to it's integer value.
// Should not be called on characters which are not digits.
func toDigit(char string) uint {
    return digitMap[char]
}

// isDigit returns true if the given character is an integer
// digit, else false.
func isDigit(char string) bool {
    _, ok := digitMap[char]
    return ok
}

// digitMap maintains a mapping of string characters to their
// integer values.
var digitMap = map[string]uint{
    "0": 0,
    "1": 1,
    "2": 2,
    "3": 3,
    "4": 4,
    "5": 5,
    "6": 6,
    "7": 7,
    "8": 8,
    "9": 9,
}

// isSign returns true if the given character is a sign character,
// else false.
func isSign(char string) bool {
    _, ok := signMap[char]
    return ok
}

// toSignScalar returns the scalar multiplier that should be used
// for a sign character.
func toSignScalar(char string) int8 {
    return signMap[char]
}

// signMap maintains a mapping from a sign character to it's scalar
// multiplier.
var signMap = map[string]int8 {
    "+": 1,
    "-": -1,
}

// isWhiteSpace returns true if the given character is whitespace.
func isWhiteSpace(char string) bool {
    return char == " "
}