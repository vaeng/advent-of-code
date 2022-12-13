defmodule AdventOfCode.Day13 do
  def part1(args) do
    args
    |> convertInputStringToLists()
    |> Enum.chunk_every(2)
    |> Enum.map(fn [left, right] -> compareLists(left, right) end)
    |> Enum.with_index(1)
    |> Enum.filter(&(elem(&1, 0) == :correct))
    |> Keyword.values()
    |> Enum.sum()
  end

  def part2(args) do
    sorted = args
    |> convertInputStringToLists()
    |> (then &([[[2]], [[6]] |&1]))
    |> Enum.sort(&compareLists(&1, &2) == :correct)

    (Enum.find_index(sorted, &(&1 == [[2]])) + 1) * (Enum.find_index(sorted, &(&1 == [[6]])) + 1)
  end

  def convertInputStringToLists(args) do
    args |> String.trim() |> (then &( "[" <> &1 <> "]")) |> String.replace("\n\n", ",") |> String.replace("\n", ",") |> Code.eval_string() |> elem(0)
  end

  def compareLists(left, right) do
    cond do
      left == right -> :continue
      right == [] -> :wrong
      left == [] -> :correct
      true ->
        [lh | lt] = left
        [rh | rt] = right
        cond do
          is_integer(lh) && is_integer(rh) && lh < rh -> :correct
          is_integer(lh) && is_integer(rh) && lh > rh -> :wrong
          is_integer(lh) && is_integer(rh) && lh == rh -> compareLists(lt, rt)

          is_list(lh) && is_integer(rh) -> compareLists([lh | lt], [[rh] | rt])
          is_integer(lh) && is_list(rh) -> compareLists([[lh] | lt], [rh | rt])

          # both lh and rh must be unequal lists:
          true -> case compareLists(lh, rh) do
                      :wrong -> :wrong
                      :correct -> :correct
                      :continue -> compareLists(lt, rt)
                  end
        end
    end
  end


end
