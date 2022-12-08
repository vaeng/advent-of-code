defmodule AdventOfCode.Day04 do
  def part1(args) do
    args|> String.split() |> Enum.map(&String.split(&1, ~r{[\-\,]})) |> Enum.map(&Enum.map(&1, fn x -> String.to_integer(x) end)) |> Enum.map(fn [a,b,c,d] -> (a >= c and b <= d) or (c >= a and d <= b)  end) |> Enum.count(fn x -> x end)
  end

  def part2(args) do
    args
    |> String.split()
    |> Enum.map(&String.split(&1, ~r{[\-\,]}))
    |> Enum.map(&Enum.map(&1, fn x -> String.to_integer(x) end))
    |> Enum.map(fn [a,b,c,d] -> (a >= c and a <= d) or (b >= c and b <= d) or (c >= a and c <= b) or (d >= a and d <= b)
   end)
    |> Enum.count(fn x -> x end)

  end
end
