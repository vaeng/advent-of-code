defmodule AdventOfCode.Day02 do
  def part1(args) do
    args
    |> String.split()
    |> Enum.map(&wrappingPaper/1)
    |> Enum.sum()
  end

  defp wrappingPaper(str), do:
    str
    |> String.split("x")
    |> Enum.map(&String.to_integer/1)
    |> Enum.sort()
    |> then (fn [x, y, z] -> (2*(x*y + y*z + z*x) + x*y) end)

  def part2(args) do
    args
    |> String.split()
    |> Enum.map(&nneededRibbon/1)
    |> Enum.sum()
  end

  defp nneededRibbon(str), do:
  str
  |> String.split("x")
  |> Enum.map(&String.to_integer/1)
  |> Enum.sort()
  |> then (fn [x, y, z] -> 2*(x+y) + x*y*z end)
end
