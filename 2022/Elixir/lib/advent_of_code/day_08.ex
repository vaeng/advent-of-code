defmodule AdventOfCode.Day08 do
  import AdventOfCode.Util

  def part1(args) do
    args |> inputToNumGrid() |> countVisibilityOnGrid()
  end

  def part2(args) do
    args |> inputToNumGrid() |> getMaxScenicScoreOnGrid()
  end

  def inputToNumGrid(input) do
    input
    |> String.trim()
    |> String.split("\n")
    |> Enum.map(
      &(String.codepoints(&1)
        |> Enum.map(fn x -> String.to_integer(x) end))
    )
  end

  @spec getScenicScore([list], integer, integer) :: number
  def getScenicScore(grid, x, y) do
    height = grid |> Enum.at(y) |> Enum.at(x)
    allWest = getAllWest(grid, x, y) |> countVisibleTrees(height)
    allEast = getAllEast(grid, x, y) |> countVisibleTrees(height)
    allSouth = getAllSouth(grid, x, y) |> countVisibleTrees(height)
    allNorth = getAllNorth(grid, x, y) |> countVisibleTrees(height)

    Enum.product([allWest, allEast, allSouth, allNorth])
  end

  def countVisibleTrees(list, height) do
    cansee = Enum.take_while(list, &(&1 < height)) |> length()
    if cansee < length(list), do: cansee + 1, else: cansee
  end

  def getMaxScenicScoreOnGrid(grid) do
    yMax = length(grid) - 1
    xMax = length(Enum.at(grid, 0)) - 1

    0..yMax
    |> Enum.map(fn y -> Enum.map(0..xMax, &[&1, y]) end)
    |> Enum.map(&Enum.map(&1, fn [x, y] -> AdventOfCode.Day08.getScenicScore(grid, x, y) end))
    |> List.flatten()
    |> Enum.max()
  end

  def countVisibilityOnGrid(grid) do
    yMax = length(grid) - 1
    xMax = length(Enum.at(grid, 0)) - 1

    0..yMax
    |> Enum.map(fn y -> Enum.map(0..xMax, &[y, &1]) end)
    |> Enum.map(&Enum.map(&1, fn [x, y] -> AdventOfCode.Day08.isVisible?(grid, x, y) end))
    |> List.flatten()
    |> Enum.count(&(&1 == true))
  end

  def isVisible?(grid, x, y) do
    height = grid |> Enum.at(y) |> Enum.at(x)
    allWest = getAllWest(grid, x, y) |> Enum.all?(&(&1 < height))
    allEast = getAllEast(grid, x, y) |> Enum.all?(&(&1 < height))
    allSouth = getAllSouth(grid, x, y) |> Enum.all?(&(&1 < height))
    allNorth = getAllNorth(grid, x, y) |> Enum.all?(&(&1 < height))

    Enum.any?([allWest, allEast, allSouth, allNorth])
  end

  @doc """
  Zero based access. Map origin is upper left corner at 0,0
  """
  def getAllWest(grid, x, y) do
    grid |> Enum.fetch!(y) |> Enum.take(x) |> Enum.reverse()
  end

  def getAllEast(grid, x, y) do
    grid |> Enum.fetch!(y) |> Enum.drop(x + 1)
  end

  def getAllSouth(grid, x, y) do
    grid |> transpose() |> Enum.fetch!(x) |> Enum.drop(y + 1)
  end

  def getAllNorth(grid, x, y) do
    grid |> transpose() |> Enum.fetch!(x) |> Enum.take(y) |> Enum.reverse()
  end
end
