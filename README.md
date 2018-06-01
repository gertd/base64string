# base64string.base64String

Helper package which automaticly base64 encodes string when marshalled to JSON and decodes when unmarshalled.
Useful for making sure marshal data is guaranteed to be encoded.

NOTE: current implementation only supports JSON encoding (implements MarshalJSON and UnmarshalJSON)