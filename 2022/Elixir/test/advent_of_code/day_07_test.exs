defmodule AdventOfCode.Day07Test do
  use ExUnit.Case

  import AdventOfCode.Day07

  @tag :skip
  test "part1" do
    input = """
    $ cd /
    $ ls
    dir a
    14848514 b.txt
    8504156 c.dat
    dir d
    $ cd a
    $ ls
    dir e
    29116 f
    2557 g
    62596 h.lst
    $ cd e
    $ ls
    584 i
    $ cd ..
    $ cd ..
    $ cd d
    $ ls
    4060174 j
    8033020 d.log
    5626152 d.ext
    7214296 k
    """

    input = String.trim(input)

    result = part1(input)

    assert result == 95437
  end

  @tag :skip
  test "part2" do
    input = nil
    result = part2(input)

    assert result
  end

  @tag :skip
  test "getSizes" do
    emptyMap = %{}
    assert getSize(emptyMap) == 0

    # nestedEmptyMap = %{"/" => %{}}
    # assert getSize(nestedEmptyMap) == 0

    # deeplyNestedEmptyMap = %{"/" => %{"abc" => %{}, "bcd" => %{"qed" => %{}}}}
    # assert getSize(deeplyNestedEmptyMap) == 0
  end
end
