require 'ffi'

module Sum
  extend FFI::Library
  ffi_lib './sum.so'
  attach_function :sum, [:int, :int], :int
end

p Sum.sum(1, 2)
