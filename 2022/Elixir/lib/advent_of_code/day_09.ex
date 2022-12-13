defmodule AdventOfCode.Day09 do
  import AdventOfCode.Util

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
    |> Enum.join()
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
end
