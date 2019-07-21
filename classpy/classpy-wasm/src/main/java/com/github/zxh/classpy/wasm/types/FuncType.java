package com.github.zxh.classpy.wasm.types;

import com.github.zxh.classpy.common.FilePart;
import com.github.zxh.classpy.wasm.WasmBinPart;
import com.github.zxh.classpy.wasm.WasmBinFile;

import java.util.stream.Collectors;

public class FuncType extends WasmBinPart {

    {
        _byte(null, (byte) 0x60);
        vector("parameters", ValType::new);
        vector("results", ValType::new);
    }

    @Override
    protected void postRead(WasmBinFile wasm) {
        String params = get("parameters").getParts().stream()
                .skip(1)
                .map(FilePart::getDesc)
                .collect(Collectors.joining(",", "(", ")"));
        String results = get("results").getParts().stream()
                .skip(1)
                .map(FilePart::getDesc)
                .collect(Collectors.joining(",", "(", ")"));
        setDesc(params + "->" + results);
    }

}
