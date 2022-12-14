defmodule AdventOfCode.Day14 do
  import AdventOfCode.Util

  def part1(args) do
    finalMap = parseInputToRockMap(args) |> dropSandUntilStable()
    renderMap(finalMap)
    finalMap |> Map.values() |> Enum.count(&(&1 == :sand))
  end

  def part2(args) do
    finalMap = parseInputToRockMap(args) |> dropSandUntilFull()
    renderMap(finalMap)
    sandcorns = finalMap |> Map.values() |> Enum.count(&(&1 == :sand))
    sandcorns + 1
  end

  def parseInputToRockMap(args) do
    flatRockPaths =
      args
      |> String.trim()
      |> String.split("\n")
      |> Enum.map(
        &(String.split(&1, [" -> ", ","])
          |> Enum.map(fn s -> String.to_integer(s) end))
      )
      |> Enum.map(&Enum.chunk_every(&1, 4, 2, :discard))
      |> List.flatten()
      |> Enum.chunk_every(2)

    {[minX, _], [maxX, _]} = flatRockPaths |> Enum.min_max_by(fn [x, _] -> x end)
    {[_, minY], [_, maxY]} = flatRockPaths |> Enum.min_max_by(fn [_, y] -> y end)

    flatRockPaths
    |> Enum.chunk_every(2)
    |> Enum.map(fn [[x1, y1], [x2, y2]] ->
      if x1 == x2,
        do: Enum.map(y1..y2, fn y -> [x1, y] end),
        else: Enum.map(x1..x2, fn x -> [x, y1] end)
    end)
    |> Enum.reduce([], &(&1 ++ &2))
    |> Map.from_keys(:rock)
    |> Map.put([500, 0], :source)
    |> Map.put(:maxX, max(maxX, 500))
    |> Map.put(:minX, min(minX, 500))
    |> Map.put(:maxY, max(maxY, 0))
    |> Map.put(:minY, min(minY, 0))
  end

  def renderMap(map) do
    (map.maxY + 1)..(map.minY - 1)
    |> Enum.reduce([], fn y, acc ->
      Enum.map((map.minX - 1)..(map.maxX + 1), fn x -> renderCell(map, [x, y]) end) ++ acc
    end)
    |> Enum.chunk_every(map.maxX - map.minX + 3)
    |> Enum.intersperse("\n")
    |> Enum.join()
    |> IO.puts()
  end

  def renderCell(map, [x, y]) do
    case Map.get(map, [x, y], :air) do
      :source -> "+"
      :rock -> "█"
      :air -> " "
      :sand -> IO.ANSI.yellow() <> "o" <> IO.ANSI.reset()
    end
  end

  def dropSand(map) do
    [restX, restY] =
      Enum.reduce_while(0..map.maxX, [500, 0], fn _x, pos -> nextSandPosition(map, pos) end)

    if restY >= map.maxY do
      Map.put(map, :halt, true)
    else
      Map.put(map, [restX, restY], :sand)
    end
  end

  def dropSandWithFloor(map) do
    [restX, restY] =
      Enum.reduce_while(0..map.maxX, [500, 0], fn _x, pos ->
        nextSandPositionWithFloor(map, pos)
      end)

    if [restX, restY] == [500, 0] do
      Map.put(map, :halt, true)
    else
      Map.put(map, [restX, restY], :sand)
    end
  end

  def dropSandUntilStable(map) do
    map = dropSand(map)
    if Map.get(map, :halt, false), do: map, else: dropSandUntilStable(map)
  end

  def nextSandPosition(map, pos) do
    [_xNow, yNow] = pos

    cond do
      yNow == map.maxY + 1 -> {:halt, pos}
      Map.get(map, getN(pos), :air) == :air -> {:cont, getN(pos)}
      Map.get(map, getNW(pos), :air) == :air -> {:cont, getNW(pos)}
      Map.get(map, getNE(pos), :air) == :air -> {:cont, getNE(pos)}
      true -> {:halt, pos}
    end
  end

  def nextSandPositionWithFloor(map, pos) do
    [_xNow, yNow] = pos

    cond do
      yNow > map.maxY -> {:halt, pos}
      Map.get(map, getN(pos), :air) == :air -> {:cont, getN(pos)}
      Map.get(map, getNW(pos), :air) == :air -> {:cont, getNW(pos)}
      Map.get(map, getNE(pos), :air) == :air -> {:cont, getNE(pos)}
      true -> {:halt, pos}
    end
  end

  def dropSandUntilFull(map) do
    map = dropSandWithFloor(map)
    if Map.get(map, :halt, false), do: map, else: dropSandUntilFull(map)
  end
end
