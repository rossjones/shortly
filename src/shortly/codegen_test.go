package main

import "testing"

func Test_Encoding(t *testing.T) {
    for x := uint64(0); x < uint64(512); x++  {
        e := base62_encode(x)
        d := base62_decode(e)
        if x != d {
            t.FailNow()
        }
    }
}


func BenchmarkBase62Encode(b *testing.B) {
    var x uint64 = 1
    for ; x < uint64(b.N + 1); x++  {
        _ = base62_encode(uint64(x))
    }
}

func BenchmarkBase62Decode(b *testing.B) {
    var x uint64 = 1
    for ; x < uint64(b.N + 1); x++  {
        _ = base62_decode(string(x))
    }
}



