defmodule AdventOfCode.Day14Test do
  use ExUnit.Case

  import AdventOfCode.Day14

  test "part1" do
    input = """
    498,4 -> 498,6 -> 496,6
    503,4 -> 502,4 -> 502,9 -> 494,9
    """

    args = String.trim(input)
    result = part1(input) == 24

    assert result
  end

  @tag :skip
  test "part2" do
    input = nil
    result = part2(input)

    assert result
  end
end
