defmodule AdventOfCode.Day01Test do
  use ExUnit.Case

  import AdventOfCode.Day01

  test "part1" do
    """
    (()) and ()() both result in floor 0.
    ((( and (()(()( both result in floor 3.
    ))((((( also results in floor 3.
    ()) and ))( both result in floor -1 (the first basement level).
    ))) and )())()) both result in floor -3.
    """

    assert  part1("(())") == 0
    assert  part1("()()") == 0
    assert  part1("(((") == 3
    assert  part1("(()(()(") == 3
    assert  part1("))(((((") == 3
    assert  part1("))(") == -1
    assert  part1("())") == -1
    assert  part1(")))") == -3
    assert  part1(")())())") == -3
  end


  test "part2" do
    """
    ) causes him to enter the basement at character position 1.
    ()()) causes him to enter the basement at character position 5.
    """
    input = nil
    assert  part2(")") == 1
    assert  part2("()())") == 5

  end
end
