class Codec:
    def encode(self, strs):
        return ''.join("{}:{}".format(len(s), s) for s in strs)

    def decode(self, s):
        decoded = []
        i = 0
        while i < len(s):
            colon = s.find(":", i)
            l = int(s[i:colon])
            decoded.append(s[colon+1:colon+1+l])
            i = colon + 1 + l
        return decoded

if __name__ == '__main__':
    # Your Codec object will be instantiated and called as such:
    codec = Codec()
    strs = ["hello", "my", "friend!", " ", "what", "say", "", "you?"]
    print codec.decode(codec.encode(strs))
    strs = []
    print codec.decode(codec.encode(strs))
    strs = ['', '']
    print codec.decode(codec.encode(strs))
    strs = ['']
    print codec.decode(codec.encode(strs))
