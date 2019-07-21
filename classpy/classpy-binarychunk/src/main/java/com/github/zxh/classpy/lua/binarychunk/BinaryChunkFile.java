package com.github.zxh.classpy.lua.binarychunk;

import com.github.zxh.classpy.lua.binarychunk.part.Function;
import com.github.zxh.classpy.lua.binarychunk.part.Header;
import com.github.zxh.classpy.lua.binarychunk.datatype.LuByte;

/**
 * Represent lua binary chunk file.
 * 
 * @see /lua/src/ldump.c#luaU_dump().
 */
public class BinaryChunkFile extends BinaryChunkPart {

    {
        add("header",        new Header());
        add("size_upvalues", new LuByte());
        add("main",        new Function());
    }

}
