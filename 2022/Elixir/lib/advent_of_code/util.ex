defmodule AdventOfCode.Util do
  @moduledoc """
  A utility module that implements functions for performing operations that
  might be used by different problem solutions.
  """

  @doc """
  Convert a newline- separated String of numbers into
  a list of integers.

  Examples

  iex> AdventOfCode.Util.stringToIntArray("1\n2\n334\n4")
  [1,2,334,4]
  """
  @spec stringToIntArray(String.t()) :: list(integer)
  def stringToIntArray(xs),
    do:
      xs
      |> String.strip()
      |> String.split("\n")
      |> Enum.map(&String.to_integer(&1))
end
