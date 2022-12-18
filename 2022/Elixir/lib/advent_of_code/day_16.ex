defmodule AdventOfCode.Day16 do
  def part1(args) do
    valveProperties = parseInput(args)
    allKeys = Map.keys(valveProperties)

    distanceMap =
      Enum.reduce(allKeys, %{}, fn x, acc ->
        Map.put(acc, x, Map.from_keys(allKeys, 1000))
      end)

    distanceMap =
      Enum.reduce(allKeys, distanceMap, fn key, acc ->
        Map.update(acc, key, %{}, &Map.merge(&1, valveProperties[key].distances))
      end)
      |> addDistancesToMap()
      |> addDistancesToMap()

    noAA = allKeys -- ["AA"]
    # noAA = Enum.filter(noAA, &(valveProperties[&1].rate != 0))

    Task.async_stream(permutate(noAA), &getResults(&1, valveProperties, distanceMap),
      ordered: false
    )
    |> Enum.max_by(fn {:ok, [_, _, score]} -> score end)
    |> then(&List.last(elem(&1, 1)))
  end

  def getResults(path, valveProperties, distanceMap) do
    Enum.reduce_while(path, [["AA"], 30, 0], fn next, [current, timeleft, score] ->
      total = timeleft - distanceMap[hd(current)][next] - 1
      neWscore = score + valveProperties[hd(current)].rate * timeleft

      if total >= 0,
        do: {:cont, [[next | current], total, neWscore]},
        else: {:halt, [current, timeleft, neWscore]}
    end)
  end

  def part2(_args) do
  end

  def parseInput(args) do
    re =
      ~r/^Valve (?<name>\w{2}) has flow rate=(?<rate>\d+); tunnels? leads? to valves? (?<tunnels>.+)$/

    args
    |> String.split("\n")
    |> Enum.map(&Regex.named_captures(re, &1))
    |> Enum.reduce(%{}, fn valve, map ->
      Map.put(map, valve["name"], %{
        rate: String.to_integer(valve["rate"]),
        tunnels: String.split(valve["tunnels"], ", "),
        distances: Map.from_keys(String.split(valve["tunnels"], ", "), 1)
      })
    end)
  end

  def addDistancesToMap(map) do
    allKeys = Map.keys(map)

    # for each key:
    Enum.reduce(allKeys, map, fn keyToBeUpdated, accMap1 ->
      # for each tunnel:
      Enum.reduce(allKeys, accMap1, fn proxytunnel, accMap2 ->
        # update distance from keyToBeUpdated to current Tunnel to all other tunnels
        Enum.reduce(allKeys, accMap2, fn endTunnel, accMap3 ->
          update_in(accMap3, [keyToBeUpdated, endTunnel], fn old ->
            # check if distance got closer
            min(
              old,
              get_in(accMap3, [keyToBeUpdated, proxytunnel]) +
                get_in(accMap3, [proxytunnel, endTunnel])
            )
          end)
        end)
      end)
    end)
  end

  @doc """
  Permutations of a list, from:
  (https://www.reddit.com/r/elixir/comments/74088r/how_can_i_improve_my_code_simple_permutation/)
  """
  def permutate([]) do
    [[]]
  end

  def permutate(list) do
    for head <- list, tail <- permutate(list -- [head]), do: [head] ++ tail
  end

  # https://github.com/alexandrubagu/elixir_stream_permutations/blob/master/lib/stream_permutations.ex
  """
  import Formulae.Combinators.H

  def generate(list) do
    length = length(list)

    ast =
      Enum.reduce(length..1, {[mapper(1, length, &var/1)], :ok}, fn i, body ->
        stream_permutation_transform_clause(i, list, body)
      end)

    {{stream, :ok}, _} = Code.eval_quoted(ast)
    stream
  end
  """
end
