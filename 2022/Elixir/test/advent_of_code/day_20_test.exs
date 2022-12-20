defmodule AdventOfCode.Day20Test do
  use ExUnit.Case

  import AdventOfCode.Day20

  test "part1" do
    args =  """
            1
            2
            -3
            3
            -2
            0
            4
            """

    args = String.trim(args)
    result = part1(args)

    assert result == 3
  end

  @tag :skip
  test "part2" do
    input = nil
    result = part2(input)

    assert result
  end
end
