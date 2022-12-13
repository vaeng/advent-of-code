defmodule AdventOfCode.Day10 do
  def part1(args) do
    cycles = calculateStateDuringCycles(args)
    interest = [20, 60, 100, 140, 180, 220]
    interest |> Enum.map(&(Enum.at(cycles, &1) * &1)) |> Enum.sum()
  end

  def part2(args) do
    cycles = calculateStateDuringCycles(args)

    cycles
    |> Enum.with_index()
    |> tl()
    |> Enum.map(fn {p, c} -> renderPixel(p, c) end)
    |> Enum.chunk_every(40)
    |> Enum.map(&Enum.join/1)
  end

  def renderPixel(p, cycle) do
    c = Integer.mod(cycle, 40)
    if p == c or p + 1 == c or p + 2 == c, do: "#", else: "."
  end

  def instructionToCycles("noop", [hd | tail]), do: [hd, hd | tail]

  def instructionToCycles(inst, [hd | tail]) do
    val = String.split(inst) |> List.last() |> String.to_integer()
    [val + hd, hd, hd | tail]
  end

  def calculateStateDuringCycles(args) do
    args
    |> String.trim()
    |> String.split("\n")
    |> Enum.reduce([1, 0], fn x, acc -> instructionToCycles(x, acc) end)
    |> Enum.reverse()
  end
end
