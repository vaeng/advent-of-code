defmodule AdventOfCode.Day09 do
  @moduledoc """
  This module is using named directions in a 2D Space.
  The center X be [0,0]. eg. NN is [0,2].
  |      |     |    |     |      |
  |:----:|:---:|:---:|:---:|:----:|
  | NWNW | NNW | NN | NNE | NENE |
  | NWW  | NW  | N  | NE  | NEE  |
  | WW   | W   | X  | E   | EE   |
  | SWW  | SW  | S  | SE  | SEE  |
  | SWSW | SSW | SS | SSE | SESE |
  """
  def part1(args) do
    args
    |> parseInput()
    |> headPositions()
    # 1
    |> nextKnotPositions()
    |> MapSet.new()
    |> MapSet.size()
  end

  def part2(args) do
    args
    |> parseInput()
    |> headPositions()
    # 1
    |> nextKnotPositions()
    # 2
    |> nextKnotPositions()
    # 3
    |> nextKnotPositions()
    # 4
    |> nextKnotPositions()
    # 5
    |> nextKnotPositions()
    # 6
    |> nextKnotPositions()
    # 7
    |> nextKnotPositions()
    # 8
    |> nextKnotPositions()
    # 9
    |> nextKnotPositions()
    |> MapSet.new()
    |> MapSet.size()
  end

  def headPositions(instructions) do
    instructions
    |> String.codepoints()
    |> Enum.scan([0, 0], fn x, acc -> AdventOfCode.Day09.updateHead(acc, x) end)
  end

  def nextKnotPositions(previousKnotPositions) do
    previousKnotPositions
    |> Enum.scan([0, 0], fn head, acc -> AdventOfCode.Day09.updateKnot(head, acc) end)
  end

  def parseInput(input) do
    input
    |> String.trim()
    |> String.split("\n")
    |> Enum.map(
      &(String.split(&1, " ")
        |> then(fn [dir, num] -> String.duplicate(dir, String.to_integer(num)) end))
    )
    |> Enum.reduce("", &(&2 <> &1))
  end

  def updateHead(h, "R"), do: getE(h)
  def updateHead(h, "L"), do: getW(h)
  def updateHead(h, "U"), do: getN(h)
  def updateHead(h, "D"), do: getS(h)

  def updateKnot(h, t) do
    cond do
      h == t -> t
      h == getN(t) -> t
      h == getNN(t) -> getN(t)
      h == getS(t) -> t
      h == getSS(t) -> getS(t)
      h == getW(t) -> t
      h == getWW(t) -> getW(t)
      h == getE(t) -> t
      h == getEE(t) -> getE(t)
      h == getSE(t) -> t
      h == getSSE(t) -> getSE(t)
      h == getSEE(t) -> getSE(t)
      h == getSESE(t) -> getSE(t)
      h == getNE(t) -> t
      h == getNNE(t) -> getNE(t)
      h == getNEE(t) -> getNE(t)
      h == getNENE(t) -> getNE(t)
      h == getSW(t) -> t
      h == getSSW(t) -> getSW(t)
      h == getSWW(t) -> getSW(t)
      h == getSWSW(t) -> getSW(t)
      h == getNW(t) -> t
      h == getNNW(t) -> getNW(t)
      h == getNWW(t) -> getNW(t)
      h == getNWNW(t) -> getNW(t)
      true -> raise "unknown skip"
    end
  end

  def getN([x, y]), do: [x, y + 1]
  def getNN([x, y]), do: [x, y + 2]
  def getS([x, y]), do: [x, y - 1]
  def getSS([x, y]), do: [x, y - 2]
  def getW([x, y]), do: [x - 1, y]
  def getWW([x, y]), do: [x - 2, y]
  def getE([x, y]), do: [x + 1, y]
  def getEE([x, y]), do: [x + 2, y]

  def getSE([x, y]), do: [x + 1, y - 1]
  def getSSE([x, y]), do: [x + 1, y - 2]
  def getSEE([x, y]), do: [x + 2, y - 1]
  def getSESE([x, y]), do: [x + 2, y - 2]

  def getNE([x, y]), do: [x + 1, y + 1]
  def getNNE([x, y]), do: [x + 1, y + 2]
  def getNEE([x, y]), do: [x + 2, y + 1]
  def getNENE([x, y]), do: [x + 2, y + 2]

  def getSW([x, y]), do: [x - 1, y - 1]
  def getSSW([x, y]), do: [x - 1, y - 2]
  def getSWW([x, y]), do: [x - 2, y - 1]
  def getSWSW([x, y]), do: [x - 2, y - 2]

  def getNW([x, y]), do: [x - 1, y + 1]
  def getNNW([x, y]), do: [x - 1, y + 2]
  def getNWW([x, y]), do: [x - 2, y + 1]
  def getNWNW([x, y]), do: [x - 2, y + 2]
end
