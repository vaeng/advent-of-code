defmodule AdventOfCode.Day19 do

  def part1(args, limit \\ 24) do
    resources = %{clay: 0, ore: 0, obsidian: 0, geodes: 0}
    workers = %{ore: 1, clay: 0, obsidian: 0, geodes: 0}
    costs = inputToCostTable(args)
    Enum.map(Map.keys(costs), fn id ->
      Enum.reduce(1..limit, [{resources,workers}],
        fn round, state ->
          Enum.flat_map(state, &progress_one_minute(&1,id,costs))
          |> Enum.sort_by(fn state -> potential(state, round, limit, costs, id) end, :desc)
          |> Enum.take(3000) end)
      |> (then &get_geode_number(&1)*id)
    end)
    |> Enum.sum()
  end

  def progress_one_minute({resources, workers}, id, costs) do
    actions = possibleActions(resources, costs, id)
    actions_paid = Enum.map(actions, &payForWorker(resources, costs, id, &1))
    updated_resources = Enum.map(actions_paid, &updateResources(&1, workers))
    updated_workers = Enum.map(actions, &updateWorkers(workers, &1))
    Enum.zip_with(updated_resources, updated_workers, &({&1, &2}))
  end

  def get_geode_number(states) do
    {topresources,_topworkers} = states
    |> Enum.sort_by(fn {resources,_workers} -> resources.geodes end, :desc)
    |> hd
    topresources.geodes
  end

  def part2(args, limit \\ 32) do
    resources = %{clay: 0, ore: 0, obsidian: 0, geodes: 0}
    workers = %{ore: 1, clay: 0, obsidian: 0, geodes: 0}
    costs = inputToCostTable(args)
    Enum.map((Map.keys(costs) |> Enum.take(3)),
     fn id -> Enum.reduce(1..limit, [{resources,workers}],
                          fn round, state -> Enum.flat_map(state, &progress_one_minute(&1,id,costs))
                                             |> Enum.sort_by(fn state -> potential(state, round, limit, costs, id) end, :desc)
                                             |> Enum.take(3000) end)
              |> (then &get_geode_number(&1))
      end)
    |> Enum.product()
  end

  def inputToCostTable(args) do
    Regex.scan(~r/\d+/, args)
    |> Enum.map(&String.to_integer(hd(&1)))
    |> Enum.chunk_every(7) |> Enum.reduce(Map.new(),
        fn [id, orecost, claycost, obscostore, obscostclay, geocostore, geocostobs], map ->
           Map.put(map, id, %{ore: %{ore: orecost}, clay: %{ore: claycost}, obsidian: %{ore: obscostore, clay: obscostclay}, geodes: %{ore: geocostore, obsidian: geocostobs} })
        end)
  end

  def updateResources(resources, workers) do
    Map.merge(resources, workers, fn _key,  old, new -> old + new end)
  end

  def payForWorker(resources, _costs, _id, :wait) do
    resources
  end
  def payForWorker(resources, costs, id, order) do
    Map.merge(resources, costs[id][order], fn _key,  old, new -> old - new end)
  end

  def possibleActions(resources, costs, id) do
    [:ore, :clay, :obsidian, :geodes]
    |> Enum.filter(&canPay?(resources, costs, id, &1))
    |> (then &(if Enum.count(&1) != 4, do: [:wait | &1], else: &1))
  end

  def canPay?(resources, costs, id, order) do
    Enum.all?(Map.merge(resources, costs[id][order], fn _key,  old, new -> old - new end), &(elem(&1,1)>=0))
  end

  def updateWorkers(workers, :wait) do
    workers
  end

  def updateWorkers(workers, order) do
    Map.update!(workers,order, &(&1 + 1))
  end

  def potential({resources, workers}, round, limit, costs, id) do
    time_left = limit - round
    potential_clay = resources.clay + workers.clay * time_left
    potential_ore = resources.ore + workers.ore * time_left
    potential_obs = resources.obsidian + workers.obsidian * time_left + div(potential_clay, costs[id].obsidian.clay)
    potential_geodes = resources.geodes + workers.geodes * time_left
    potential_geodes * 10000 + potential_obs * 100 + potential_ore + potential_clay
  end
end
