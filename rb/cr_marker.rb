# C-T2 marker: a required file that defines a custom cop emitting a marker
# offense. If require executes, the marker appears in findings.
# NOTE: the Kudelski vector (require: ./ext.rb) is DEAD — do not retest it
# as-is; the variants below use path/anchor tricks instead.
module RuboCop
  module Cop
    module Custom
      class CrMarker < Base
        MSG = "CR-MARKER-T2-7f3a9 required file executed"
        def on_new_investigation
          add_offense(processed_source.ast, message: MSG) if processed_source.ast
        end
      end
    end
  end
end
