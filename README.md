hasher
======

Hasher passes stdin to multiple hash functions and outputs checksum from each in hex and base64 format.

    > echo -en "bacon" | hasher
    hash      hex                                                              base64
    adler32   05e80204                                                         BegCBA==
    crc32     a4b40d18                                                         pLQNGA==
    crc64     3d8c377416f00000                                                 PYw3dBbwAAA=
    keccak224 1dca0dc59ffc2bb4e0c1080f8cc324faeed607fe023246741aef8369         HcoNxZ/8K7TgwQgPjMMk+u7WB/4CMkZ0Gu+DaQ==
    keccak256 21ededac5207f57ff13b263856c380f44e000a0f3e8ce86ce9f9af087de8a4c4 Ie3trFIH9X/xOyY4VsOA9E4ACg8+jOhs6fmvCH3opMQ=
    md4       62445769fb62f35ca36acffde98ae3a9                                 YkRXafti81yjas/96YrjqQ==
    md5       7813258ef8c6b632dde8cc80f6bda62f                                 eBMljvjGtjLd6MyA9r2mLw==
    ripemd160 58ff8220275e6592716d40fbe8fc19283304dc99                         WP+CICdeZZJxbUD76PwZKDME3Jk=
    sha1      8abf15bef376e0e21f1f9e9c3d74483d5018f3d5                         ir8VvvN24OIfH56cPXRIPVAY89U=
    sha256    9cca0703342e24806a9f64e08c053dca7f2cd90f10529af8ea872afb0a0c77d4 nMoHAzQuJIBqn2TgjAU9yn8s2Q8QUpr46ocq+woMd9Q=

