"""

Given two strings, check to see if they are anagrams. An anagram is when the two
strings can be written using the exact same letters (so you can just rearrange
the letters to get a different phrase or word).

For example:

"public relations" is an anagram of "crap built on lies."

"clint eastwood" is an anagram of "old west action"

Note: Ignore spaces and capitalization. So "d go" is an anagram of
"God" and "dog" and "o d g".

"""

import sys
from nose.tools import assert_equal


def anagram(s1, s2):

    # We will not do the following since the sorted() function in python has
    # O(nlog(n)) complexity, and we can do better
    # return sorted(s1.lower().replace(" ", "")) == \
    #        sorted(s2.lower().replace(" ", ""))

    s1 = s1.lower().replace(" ", "")
    s2 = s2.lower().replace(" ", "")

    # now we have 2 strings without any spaces and all lowercased

    # if the lengths do not match, then it's not an anagram anyway
    if len(s1) != len(s2):
        return False

    # this dictionary will hold the count for each character in the string
    character_count = {}

    # increment counts of every character
    for character in s1:
        if character in character_count:
            character_count[character] += 1
        else:
            character_count[character] = 1

    # decrement one for every character encountered, return False if some
    # character does not exist in the count dict.
    for character in s2:
        if character in character_count:
            character_count[character] -= 1
        else:
            return False

    # check if all the counts are 0 now, if not, not an anagram
    for character in character_count:
        if character_count[character] != 0:
            return False

    return True


class AnagramTest(object):
    def test(self, sol):
        assert_equal(sol('go go go', 'gggooo'), True)
        assert_equal(sol('abc', 'cba'), True)
        assert_equal(sol('hi man', 'hi     man'), True)
        assert_equal(sol('aabbcc', 'aabbc'), False)
        assert_equal(sol('123', '1 2'), False)
        print("ALL TEST CASES PASSED")


if __name__ == "__main__":

    try:
        print(anagram(sys.argv[1], sys.argv[2]))
    except IndexError:
        print("provide 2 strings in quotes")
        sys.exit(1)

    # run tests now
    t = AnagramTest()
    t.test(anagram)