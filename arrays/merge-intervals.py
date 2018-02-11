#!/usr/bin/python3

# Given a set of non-overlapping intervals, insert a new interval into the intervals (merge if necessary).
#
# You may assume that the intervals were initially sorted according to their start times.
#
# Example 1:
#
# Given intervals [1,3],[6,9] insert and merge [2,5] would result in [1,5],[6,9].
#
# Example 2:
#
# Given [1,2],[3,5],[6,7],[8,10],[12,16], insert and merge [4,9] would result in [1,2],[3,10],[12,16].
#
# This is because the new interval [4,9] overlaps with [3,5],[6,7],[8,10].
#
# Make sure the returned intervals are also sorted.

# Definition for an interval.
class Interval:
    def __init__(self, s=0, e=0):
        self.start = s
        self.end = e

class Solution:
    # @param intervals, a list of Intervals
    # @param new_interval, a Interval
    # @return a list of Interval
    def insert(self, intervals, new_interval):
        found = False
        found_interval_end = False
        for index, interval in enumerate(intervals):
            if new_interval.start < interval.end:
                found = True
                new_interval_index = index
                if new_interval.start < interval.start:
                    # print('\n{} is lesser than {} and lesser than {}, inserting at {}\n'.format(new_interval.start, interval.end, interval.start, new_interval_index))
                    intervals.insert(new_interval_index, Interval(new_interval.start))

                for interval2 in intervals[index+1:]:
                    if new_interval.end < interval2.end:
                        found_interval_end = True
                        if new_interval.end >= interval2.start:
                            # print('\n{} is lesser than {} but greater than {}, interval ends with {}, deleting "{} : {}"\n'.format(new_interval.end, interval2.end, interval2.start, interval2.end, interval2.start, interval2.end))
                            intervals.remove(interval2)
                            intervals[new_interval_index].end = interval2.end
                        else:
                            # print('\n{} is lesser than {} and lesser than {}, interval ends with {}\n'.format(new_interval.end, interval2.end, interval2.start, new_interval.end))
                            intervals[new_interval_index].end = new_interval.end
                        break
                    else:
                        # print('Deleting "{} : {}"'.format(interval2.start, interval2.end))
                        intervals.remove(interval2)
                if not found_interval_end:
                    intervals[new_interval_index].end = new_interval.end
                break
        if not found:
            intervals.append(new_interval)
        return intervals


def print_intervals(intervals):
    for i in intervals:
        print('{} , {}\n'.format(i.start, i.end), end="")


if __name__ == '__main__':
    s = Solution()
    input = [Interval(3, 5), Interval(8, 10)]
    new_interval = Interval(1, 12)
    print('Inserting "{} : {}" to input:'.format(new_interval.start, new_interval.end))
    print_intervals(input)
    output = s.insert(input, new_interval)
    print('Output is:')
    print_intervals(output)
