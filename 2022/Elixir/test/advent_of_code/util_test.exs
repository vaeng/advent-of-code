defmodule AdventOfCode.Day01Test do
  use ExUnit.Case

  import AdventOfCode.Util

  test "stringToIntArray" do
    # test regular
    assert stringToIntArray("1\n2\n334\n4") == [1, 2, 334, 4]
    # test additional break in the end
    assert stringToIntArray("1\n2\n334\n4\n") == [1, 2, 334, 4]
    # test negative input
    assert stringToIntArray("-1\n2\n-34\n4\n") == [-1, 2, -34, 4]
    # test  additional break in the front
    assert stringToIntArray("\n-1\n2\n-34\n4\n") == [-1, 2, -34, 4]
  end
end
