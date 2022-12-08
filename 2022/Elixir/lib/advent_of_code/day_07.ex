defmodule AdventOfCode.Day07 do
  def part1(_args) do
  end

  def part2(_args) do
  end

  @spec getSize(list) :: number
  @doc """
  Returns the size of a filetree in form of a Map.
  """
  def getSize(tree) do
    tree |> Map.to_list() |> List.flatten() |> Enum.map(&getSize/1) |> Enum.sum()
  end
end
