defmodule AdventOfCode.Day08 do
  def part1(_args) do
  end

  def part2(_args) do
  end

  def inputToNumGrid(input) do
    input
    |> String.trim()
    |> String.split("\n")
    |> Enum.map(&(String.codepoints(&1)
                   |> Enum.map(fn x -> String.to_integer(x) end)))
  end

  def setAllToUnknown(numgrid) do
    numgrid |> Enum.map(&Enum.map(&1, fn x -> [x, :unknown] end))
  end

  def initializeBoarders(tagGrid) do
    tagGrid
    |> setFirstRowVis() # top
    |> transpose()
    |> setFirstRowVis() # left
    |> transpose()
    |> setFirstRowVis() # down
    |> transpose()
    |> setFirstRowVis() # right
    |> transpose() # return in original orientation
  end

  def setFirstRowVis([firstRow | rest]) do
    visFirst = firstRow |> Enum.map( fn [x, _] -> [x, :vis] end)
    [visFirst | rest]
  end

  def transpose(rows) do
    rows |> List.zip|> Enum.map(&Tuple.to_list/1)
  end
end
