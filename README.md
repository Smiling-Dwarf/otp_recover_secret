OTP Recover Secret
==================

Context
-------

I was feedup reading it was impossible to recover your OTP secrets from Google Authenticator or similar. I created this little program to read Google Authenticator export QR code and display the secret key.

How to use
----------

Download the lastest binary or build the binary

Get the string encoded in the QR code: exemple
    otpauth-migration://offline?data=Ch4KC2Yk6WxR3xW8Rde7EglGYWtlIEdBIDEgASgBMAIKIgoPA10zkjJ0tijvit4MRqKSEglGYWtlIEdBIDIgASgBMAIQARgBIAAoz5a69vj%2F%2F%2F%2F%2FAQ%3D%3D

Pass the "data" value to the binary. You'll get the description of the record on the first line, then the readable secret on the second line:
    $ otp_recover Ch4KC2Yk6WxR3xW8Rde7EglGYWtlIEdBIDEgASgBMAIKIgoPA10zkjJ0tijvit4MRqKSEglGYWtlIEdBIDIgASgBMAIQARgBIAAoz5a69vj%2F%2F%2F%2F%2FAQ%3D%3D
    Record 0 : secret:"f$\xe9lQ\xdf\x15\xbcE׻"  name:"Fake GA 1"  algorithm:ALGORITHM_SHA1  digits:DIGIT_COUNT_SIX  type:OTP_TYPE_TOTP
    MYSOS3CR34K3YROXXM======
    Record 1 : secret:"\x03]3\x922t\xb6(\xef\x8a\xde\x0cF\xa2\x92"  name:"Fake GA 2"  algorithm:ALGORITHM_SHA1  digits:DIGIT_COUNT_SIX  type:OTP_TYPE_TOTP
    ANOTHERSOS3CR34K3YGENIUS
