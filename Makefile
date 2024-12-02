all: go_and_c chans structs

go_and_c: go_and_c_basic01 go_and_c_basic02 go_and_c_pointer01 go_and_c_array01 go_and_c_array02

go_and_c_basic01:
	$(MAKE) -C go_and_c/basic01

go_and_c_basic02:
	$(MAKE) -C go_and_c/basic02

go_and_c_pointer01:
	$(MAKE) -C go_and_c/pointer01

go_and_c_array01:
	$(MAKE) -C go_and_c/array01

go_and_c_array02:
	$(MAKE) -C go_and_c/array02

chans: chans_chan01

chans_chan01:
	$(MAKE) -C chans/chan01

structs: structs_interface_to_slice01

structs_interface_to_slice01:
	$(MAKE) -C structs/interface_to_slice01

clean:
	find $(BUILD_DIR) -type f ! -name ".*" -exec rm -f {} +

.PHONY: go_and_c chans structs