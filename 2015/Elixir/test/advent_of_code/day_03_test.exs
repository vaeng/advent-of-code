defmodule AdventOfCode.Day03Test do
  use ExUnit.Case

  import AdventOfCode.Day03

  test "part1" do

    assert part1(">") == 2
    assert part1("^>v<") == 4
    assert part1("^v^v^v^v^v") == 2
  end

  test "part2" do
    assert part1(">") == 3
    assert part1("^>v<") == 3
    assert part1("^v^v^v^v^v") == 11
  end
end
