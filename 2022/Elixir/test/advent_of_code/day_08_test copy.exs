defmodule AdventOfCode.Day08Test do
  use ExUnit.Case

  import AdventOfCode.Day08

  test "the number of characters of code for string literals" do
    input1 = "\"\""
    input2 = "\"abc\""
    input3 = "\"aaa\"aaa\""
    input4 = "\"\x27\""
    allInput = input1 <> input2 <> input3 <> input4

    assert numberOfCodeChars(input1) == 2
    assert numberOfCodeChars(input2) == 5
    assert numberOfCodeChars(input3) == 10
    assert numberOfCodeChars(input4) == 6
    assert numberOfCodeChars(allInput) == 23
  end

  test "the number of characters in memory for the values of the strings" do
    input1 = "\"\""
    input2 = "\"abc\""
    input3 = "\"aaa\"aaa\""
    input4 = "\"\x27\""
    allInput = input1 <> input2 <> input3 <> input4

    assert numberOfMemoryChars(input1) == 2
    assert numberOfMemoryChars(input2) == 5
    assert numberOfMemoryChars(input3) == 10
    assert numberOfMemoryChars(input4) == 7
    assert numberOfMemoryChars(allInput) == 23
  end

  @tag :skip
  test "part1" do
    input1 = "\"\""
    input2 = "\"abc\""
    input3 = "\"aaa\"aaa\""
    input4 = "\"\x27\""
    allInput = input1 <> input2 <> input3 <> input4

    assert part1(allInput) == 12
  end

  @tag :skip
  test "part2" do
    input = nil
    result = part2(input)

    assert result
  end
end
