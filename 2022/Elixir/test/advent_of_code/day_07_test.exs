defmodule AdventOfCode.Day07Test do
  use ExUnit.Case

  import AdventOfCode.Day07

  @tag :skip
  test "part1" do
    input = nil
    result = part1(input)

    assert result
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
