defmodule MatchingBrackets do
  @doc """
  Checks that all the brackets and braces in the string are matched correctly, and nested correctly
  """
  @spec check_brackets(String.t()) :: boolean
  def check_brackets(str) do
    String.graphemes(str)
    |> check_brackets([])
  end


  defp check_brackets([], []), do: true
  defp check_brackets([], _), do: false
  defp check_brackets([ch | t], stack) do
    case ch do
      "[" -> check_brackets(t, ["[" | stack])
      "]" -> check_brackets(ch, t, stack)
      "{" -> check_brackets(t, ["{" | stack])
      "}" -> check_brackets(ch, t, stack)
      "(" -> check_brackets(t, ["(" | stack])
      ")" -> check_brackets(ch, t, stack)
      _ -> check_brackets(t, stack)
    end
  end

  defp check_brackets(_ch, _, []), do: false
  defp check_brackets(ch, t, [s | stack_tail]) do
    case ch do
      "]" -> if s == "[", do: check_brackets(t, stack_tail), else: false
      "}" -> if s == "{", do: check_brackets(t, stack_tail), else: false
      ")" -> if s == "(", do: check_brackets(t, stack_tail), else: false
    end
  end  
end
