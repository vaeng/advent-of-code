defmodule AdventOfCode.Day09 do
  def part1(args) do
    instructions = parseInput(args)

    headPositions =
      instructions
      |> String.codepoints()
      |> Enum.scan([0, 0], fn x, acc -> AdventOfCode.Day09.updateHead(acc, x) end)

    headPositions
    |> Enum.scan([0, 0], fn head, acc -> AdventOfCode.Day09.updateTail(head, acc) end)
    |> MapSet.new()
    |> MapSet.size()
  end

  def part2(_args) do
  end

  def parseInput(input) do
    input
    |> String.trim()
    |> String.split("\n")
    |> Enum.map(
      &(String.split(&1, " ")
        |> then(fn [dir, num] -> String.duplicate(dir, String.to_integer(num)) end))
    )
    |> Enum.reduce("", fn x, y -> y <> x end)
  end

  def updateHead([x, y], "R"), do: [x + 1, y]
  def updateHead([x, y], "L"), do: [x - 1, y]
  def updateHead([x, y], "U"), do: [x, y + 1]
  def updateHead([x, y], "D"), do: [x, y - 1]

  def updateTail([xh, yh], [xt, yt]) do
    newXt =
      cond do
        yh == yt and xh - xt > 1 -> xt + 1
        yh == yt and xh - xt < -1 -> xt - 1
        yh > yt + 1 and xh > xt -> xt + 1
        yh < yt -1  and xh < xt -> xt - 1
        true -> xt
      end

    newYt =
      cond do
        xh == xt and yh - yt > 1 -> yt + 1
        xh == xt and yh - yt < -1 -> yt - 1
        xh > xt + 1 and yh > yt -> yt + 1
        xh < xt - 1 and yh < yt -> yt - 1
        true -> yt
      end

    [newXt, newYt]
  end
end
