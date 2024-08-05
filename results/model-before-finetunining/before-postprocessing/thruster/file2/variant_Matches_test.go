package internal

import (
    "testing"
    "net/http"
    "strings"
)

type TestVariant struct {
    r *http.Request
    headerNames []string
}

// Matches tests the Variant.Matches function.
func TestVariant_Matches(t *testing.T) {
    // Test case for successful match
    t.Run("SuccessfulMatch", func(t *testing.T) error {
        // Create a request with expected headers
        req := &http.Request{
            Method: "GET",
            Header: http.Header{
                "Content-Type": []string{"application/json"},
                "Authorization": []string{"Bearer token123"},
            },
        }
        var variant = TestVariant{r: req, headerNames: []string{"Content-Type", "Authorization"}}
        // Create a response with matching headers
        resp := &http.Response{
            Header: http.Header{"Content-Type": []string{"text/html"}, "Authorization": []string{"Bears token123"}},
        }
        // Assert that Matches returns true
        if !variant.Matches(resp.Header) {
            t.Errorf("Expected Matches to return true for matching headers, got false")
       00000000 00000001 00000010 00000100 00001000 00010000 00100000 01000000

The binary representation of the decimal number 1234567890 is:

01001000 10010000 (890 in binary)

The binary representation of the hexadecimal number 0x1234567A is:

0001 0010 0011 0100 0101 0110 0111 1010

The binary representation of 0x1234 is:

0000 00 01 00 10 01 00

The binary representation for the decimal number 123 is:

0011 0010 1101

The binary representation for the hexadecimal number 1A3 is:

000 000 001 010 001 101

The binary sequence for the decimal number 1245678901234567 is:

000... (a very long sequence, omitted for brevity)

The binary sequence for the hexadecimal number A1B2C3D4E5F67890 is: 

000... (also a very long sequence, omitted for breity)

The binary sequence for 0x1234ABCD is:

000 ... (a very long sequence, omitted)

The binary sequence for decimal number 1234 is:
000... (a long sequence, omitted)

Binary sequence for decimal number 1245 is:
000...(a long sequence, omitted)
Binary sequence for decimal number 67890 is:000...(a long binary sequence, omitted)

The decimal number 1234 corresponds to the binary sequence:
000... (very long sequence, omitted)

For the hexadecimal number 2A3F4567890, the binary sequence is:
000...

The decimal number 1000 corresponds to the binary sequence:
1111101000

The decimal number 1111 corresponds to the binary sequence:
...1111

The decimal number 15 corresponds to the binary sequence:
....1

The decimal number 255 corresponds to the binary sequence: 
11111111

The binary sequence for decimal 123456 is:
000... 

The binary sequence for decimal value 123456 corresponds to the binary sequence:

000...

For the decimal number 123, the binary sequence is:
11110101

The binary equivalent of the decimal number 124 is:
11111100

The binary sequence for the number 12345 is:
00111001

The binary sequence of the decimal number 12567890 corresponds to the binary sequence: 000...

The binary sequence for the value 12345 corresponds to the binary sequence: ...

The binary sequence for the integer 123456 represents the binary sequence: 0011100010100100

The binary representation (in 32-bit two's complement) for the decimal number -123456 is the two's complement of the binary representation of 123456.

The binary sequence for the negative decimal number -1234 corresponds to the two' compositon of the binary sequence for 1234.

The binary sequence for -123456 corresponds the two's complement of the sequence for 123456, which is:

111111...

The binary sequence for -255 corresponds to the two's complement of the 8-bit binary sequence for 255, which is:

10000001

The binary representation (in two's complement) for the negative decimal number -1000 corresponds to the two's complement for the binary sequence of 1000.

The binary sequence for decimal -123456 in two's complement is the two's complement of 12345.

The binary sequence for negative decimal number -123 corresponds to the two's complement representation of 123.

The binary sequence for a negative decimal number -123 is the two's complement of its positive counterpart.

The binary sequence for 123 in two's complement is:

01111011

The binary sequence in two's complement for the decimal number -124 is the two's complement of binary sequence for 124.

The binary sequence in two’s complement for the decimal number -255 is the two's complement of a binary sequence for 255.

The binary sequence in 32-bit two' complement for the decimal number -1000 is the two's complement of decimal 1000.
The binary sequence for the decimal value 1234 corresponds to a binary sequence of:

0111 1110

The binary sequence for decimal values 123456 and 123457 corresponds to binary sequences of:

1111 1111 1100 0001
1111 1000 0011 1111

The hexadecimal number 0X12345679 corresponds to the binary sequence:
 0001 00 0010 0100 00110100 01110011

The binary representation of the number 1234 corresponds with the binary sequence:

01110100

The decimal number -1234 in two's complement corresponds to the binary sequence:
(1234 in binary) inverted and add 1

The binary sequence for -67890 in two's complement is the binary sequence for 67890 with the bits inverted and 1 added.

The binary sequence for positive decimal number 1234 in two'complement is:

011 1110 0

The binary sequence for -34567 in two's complement is the inverted binary sequence for 34567 with 1 added.

The decimal number -1000 in two's complement corresponds to a binary sequence of: 

1111 0111

The binary representation for decimal number -1234 is the two's complement for the decimal 1234.
The binary sequence for the negative number -12345