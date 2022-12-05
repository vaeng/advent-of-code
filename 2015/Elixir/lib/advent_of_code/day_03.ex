defmodule AdventOfCode.Day03 do
  def part1(args) do
    args
    |> String.to_charlist()
    |> Enum.reduce([[0,0], MapSet.new([[0,0]])], &addPosition(&1, &2))
    |> List.last()
    |> MapSet.size()

  end

  @spec addPosition(integer, MapSet.t()) :: MapSet.t()
  defp addPosition(?^, [[x,y],map]), do: [[x, y+1], MapSet.put(map, [x, y+1])]
  defp addPosition(?<, [[x,y],map]), do: [[x-1, y], MapSet.put(map, [x-1, y])]
  defp addPosition(?>, [[x,y],map]), do: [[x+1, y], MapSet.put(map, [x+1, y])]
  defp addPosition(?v, [[x,y],map]), do: [[x, y-1], MapSet.put(map, [x, y-1])]

  defp addPosition(?^, [:santa, [a,b], [x,y],map]), do: [:robo, [a,b], [x, y+1], MapSet.put(map, [x, y+1])]
  defp addPosition(?<, [:santa, [a,b], [x,y],map]), do: [:robo, [a,b], [x-1, y], MapSet.put(map, [x-1, y])]
  defp addPosition(?>, [:santa, [a,b], [x,y],map]), do: [:robo, [a,b], [x+1, y], MapSet.put(map, [x+1, y])]
  defp addPosition(?v, [:santa, [a,b], [x,y],map]), do: [:robo, [a,b], [x, y-1], MapSet.put(map, [x, y-1])]

  defp addPosition(?^, [:robo, [a,b], [x,y],map]), do: [:santa, [a,b+1], [x, y], MapSet.put(map, [a,b+1])]
  defp addPosition(?<, [:robo, [a,b], [x,y],map]), do: [:santa, [a-1,b], [x, y], MapSet.put(map, [a-1,b])]
  defp addPosition(?>, [:robo, [a,b], [x,y],map]), do: [:santa, [a+1,b], [x, y], MapSet.put(map, [a+1,b])]
  defp addPosition(?v, [:robo, [a,b], [x,y],map]), do: [:santa, [a,b-1], [x, y], MapSet.put(map, [a,b-1])]

  def part2(args) do
    args
    |> String.to_charlist()
    |> Enum.reduce([:santa, [0,0], [0,0], MapSet.new([[0,0]])], &addPosition(&1, &2))
    |> List.last()
    |> MapSet.size()

  end



end
