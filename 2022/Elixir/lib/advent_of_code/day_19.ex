defmodule AdventOfCode.Day19 do



  def part1(args) do
    resources = %{clay: 0, ore: 0, obsidian: 0, geodes: 0}
    workers = %{ore: 1, clay: 0, obsidian: 0, geodes: 0}
    costs = inputToCostTable(args)

  end

  def part2(_args) do
  end

  def inputToCostTable(args) do
    Regex.scan(~r/\d+/, args) |> Enum.map(&String.to_integer(hd(&1))) |> Enum.chunk_every(7) |> Enum.reduce(Map.new(), fn [id, orecost, claycost, obscostore, obscostclay, geocostore, geocostobs], map -> Map.put(map, id, %{ore: %{ore: orecost}, clay: %{ore: claycost}, obsidian: %{ore: obscostore, clay: obscostclay}, geodes: %{ore: geocostore, obsidian: geocostobs} }) end)
  end

  def updateResources(resources, workers) do
    Map.merge(resources, workers, fn _key,  old, new -> old + new end)
  end

  def payForWorker(resources, costs, id, order) do
    Map.merge(resources, costs[id][order], fn _key,  old, new -> old - new end)
  end

  def possibleActions(resources, costs, id) do
    [:ore, :clay, :obsidian, :geodes]
    |> Enum.filter(&canPay?(resources, costs, id, &1))
  end

  def canPay?(resources, costs, id, order) do
    Enum.all?(Map.merge(resources, costs[id][order], fn _key,  old, new -> old - new end), &(elem(&1,1)>=0))
  end

end
