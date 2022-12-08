defmodule AdventOfCode.Day08Test do
  use ExUnit.Case

  import AdventOfCode.Day08

  @tag :skip
  test "part1" do
    input =
      """
      30373
      25512
      65332
      33549
      35390
      """

    assert part1(allInput) == 12
  end

  @tag :skip
  test "part2" do
    input = nil
    result = part2(input)

    assert result
  end
end
