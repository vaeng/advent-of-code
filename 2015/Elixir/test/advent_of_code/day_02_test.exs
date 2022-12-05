defmodule AdventOfCode.Day02Test do
  use ExUnit.Case

  import AdventOfCode.Day02

  test "part1" do

    assert part1("2x3x4") == 58
    assert part1("1x1x10") == 43
  end

  test "part2" do

    assert part2("2x3x4") == 34
    assert part2("1x1x10") == 14
  end
end
