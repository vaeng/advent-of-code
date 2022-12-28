defmodule AdventOfCode.Day21 do
  def part1(args) do
    getResult(parseInput(args), "root")
  end

  def parseInput(args) do
    args
    |> String.trim()
    |> String.split("\n")
    |> Enum.reduce(Map.new(), fn line, map ->
        [node, rawinstr | _] = String.split(line, ": ")
        instr = if Regex.match?(~r/\d+/, rawinstr), do: String.to_integer(rawinstr), else: rawinstr
        Map.put(map, node, instr)
      end)
  end

  def getResult(map, key) do
    instr = Map.fetch!(map, key)
    if is_integer(instr) do
      instr
    else
      resolveInstruction(map, String.split(instr))
    end
  end

  def resolveInstruction(map, [key1, op, key2]) do
    oepration = case op do
      "+" -> &(&1 + &2)
      "-" -> &(&1 - &2)
      "*" -> &(&1 * &2)
      "/" -> &(&1 / &2)
    end
    oepration.(getResult(map, key1),getResult(map, key2))
  end

  def part2(args) do
    fitSides(parseInput(args), "root")
  end

  def fitSides(map, "root") do
    [leftKey, _, rightKey | _] = Map.fetch!(map, "root") |> String.split()

    case [fitSides(map, leftKey), fitSides(map, rightKey)] do
      [{:ok, resultLeft}, {:nok}] -> findFitting(map, rightKey, resultLeft)
      [{:nok}, {:ok, resultRight}] -> findFitting(map, leftKey, resultRight)
    end
  end

  def fitSides(_map, "humn") do
    {:nok}
  end

  def fitSides(map, key) do
    instr = Map.fetch!(map, key)
    if is_integer(instr) do
      {:ok, instr}
    else
      [key1, op, key2] = String.split(instr)
      operation = case op do
        "+" -> &(&1 + &2)
        "-" -> &(&1 - &2)
        "*" -> &(&1 * &2)
        "/" -> &(&1 / &2)
      end
      case [fitSides(map, key1), fitSides(map, key2)] do
        [{:ok, int1}, {:ok, int2}] -> {:ok, operation.(int1, int2)}
        _ -> {:nok}
      end
    end
  end

  def findFitting(map, key, fitting) do
    if key == "humn" do
      fitting
    else
      [leftKey, op, rightKey] = Map.fetch!(map, key) |> String.split()
      reversedOperationFindLeft = case op do
        "+" -> &(&1 - &2)
        "-" -> &(&1 + &2)
        "*" -> &(&1 / &2)
        "/" -> &(&1 * &2)
      end
      reversedOperationFindRight = case op do
        "+" -> &(&1 - &2)
        "-" -> &(&2 - &1)
        "*" -> &(&1 / &2)
        "/" -> &(&2 / &1)
      end
      case [fitSides(map, leftKey), fitSides(map, rightKey)] do
        [{:ok, resultLeft}, {:nok}] -> findFitting(map, rightKey, reversedOperationFindRight.(fitting, resultLeft))
        [{:nok}, {:ok, resultRight}] -> findFitting(map, leftKey, reversedOperationFindLeft.(fitting, resultRight))
      end
    end
  end


end
