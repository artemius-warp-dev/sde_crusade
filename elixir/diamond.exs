defmodule Diamond do
  @doc """
  Given a letter, it prints a diamond starting with 'A',
  with the supplied letter at the widest point.
  """
  @spec build_shape(char) :: String.t()
  def build_shape(letter) do
    build_shape(?A, letter, acc, ?A - letter, ?A - letter)
  end

  # defp build_shape(end_letter, end_letter, acc, 0) do
    
  #   build_shape()
  # end

  # defp build_shape(letter, end_letter,  acc, dots_length) do
  #   
  #   
  # end

  defp build_shape(?A, end_letter, acc, dots_amount, dots_reminder) do
    updated_acc =  ["#{String.duplicate(".", dots_amount)}#{[?A]}#{String.duplicate(".", dots_amount)}" | acc]
    build_shape(?A + 1, end_letter, updated_acc, dots_amount, dots_reminder - 1)
  end


  defp build_shape(end_letter, end_letter, acc, dots_amount, 0) do
    updated_acc = ["#{[letter]}#{String.duplicate(".", dots-amount)}#{[letter]}" | acc]
    build_shape(end_letter, ?A, updated_acc, dots_amount, dots_remainder)
  end

  defp build_shape(letter, end_letter, acc, dots_amount, dots_reminder) do
    #···B·B···
    updated_acc =  ["#{String.duplicate(".", dots_reminder)}#{[letter]}#{String.duplicate(".", dots-amount - dots_reminder)}#{[letter]}#{String.duplicate(".", dots_amount)}" | acc]
    build_shape(letter + 1, end_letter, updated_acc, dots_amount, dots_reminder - 1)
  end


  
  
end


Diamond.build_shape(?A)
