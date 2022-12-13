defmodule AdventOfCode.Day12 do
  import AdventOfCode.Util

  def part1(args) do
    heightMap = getHeightmap(args)
    [{stop, _}, {start, _} | _] = heightMap |> Map.to_list() |> List.keysort(1)

    # clean start and stop:
    heightMap = heightMap |> Map.put(start, 0) |> Map.put(stop, ?z - ?a)

    queue = [{start, 0, heuristic(start, stop)}]
    updateUntilStop(queue, heightMap, stop)
  end

  def part2(args) do
    heightMap = getHeightmap(args)
    [{stop, _}, {start, _} | _] = heightMap |> Map.to_list() |> List.keysort(1)

    # clean start and stop:
    heightMap = heightMap |> Map.put(start, 0) |> Map.put(stop, ?z - ?a)

    {[xMax, _], _} = heightMap |> Enum.max()

    startingPoints = 1..xMax |> Enum.map(&[&1, 0])

    startingPoints
    |> Enum.map(&updateUntilStop([{&1, 0, heuristic(start, stop)}], heightMap, stop))
    |> Enum.min()
  end

  @spec heuristic([number, ...], [number, ...]) :: number
  def heuristic([cx, cy], [gx, gy]) do
    abs(gx - cx) + abs(gy - cy)
  end

  def getHeightmap(args) do
    args
    |> String.split("\n")
    |> Enum.with_index(fn row, rix ->
      String.codepoints(row)
      |> Enum.with_index(fn <<v::utf8>>, cix -> %{[rix, cix] => v - ?a} end)
    end)
    |> List.flatten()
    |> Enum.reduce(%{}, &Map.merge(&1, &2))
  end

  def getPossibleMoves(heightMap, currentPosition, step) do
    [getN(currentPosition), getS(currentPosition), getW(currentPosition), getE(currentPosition)]
    |> Enum.filter(&Map.has_key?(heightMap, &1))
    |> Enum.filter(&(Map.fetch!(heightMap, &1) <= 1 + Map.fetch!(heightMap, currentPosition)))
    |> Enum.map(&{&1, step + 1})
  end

  def updateQueue(queue, heightMap, stop) do
    posMoves =
      queue
      |> Enum.map(&getPossibleMoves(heightMap, elem(&1, 0), elem(&1, 1)))
      |> Enum.flat_map(& &1)

    if List.keymember?(posMoves, stop, 0) do
      [{stop, elem(List.keyfind!(posMoves, stop, 0), 1)}]
    else
      posMoves
      |> Enum.reduce(
        queue,
        fn {move, step}, accQueue ->
          if List.keymember?(accQueue, move, 0),
            do: accQueue,
            else:
              List.keystore(
                accQueue,
                move,
                0,
                {move, step, heuristic(move, stop)}
              )
        end
      )
    end
  end

  def updateUntilStop(queue, heightMap, stop) do
    if List.keymember?(queue, stop, 0) do
      elem(List.keyfind(queue, stop, 0), 1)
    else
      nextQueue = updateQueue(queue, heightMap, stop)
      updateUntilStop(nextQueue, heightMap, stop)
    end
  end
end
