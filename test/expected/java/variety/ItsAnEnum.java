/**
 * Autogenerated by Frugal Compiler (1.11.0)
 * DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING
 *
 * @generated
 */
package variety.java;

import java.util.Map;
import java.util.HashMap;
import org.apache.thrift.TEnum;

public enum ItsAnEnum implements org.apache.thrift.TEnum {
	FIRST(2),
	SECOND(3),
	THIRD(4);

	private final int value;

	private ItsAnEnum(int value) {
		this.value = value;
	}

	public int getValue() {
		return value;
	}

	public static ItsAnEnum findByValue(int value) {
		switch (value) {
			case 2:
				return FIRST;
			case 3:
				return SECOND;
			case 4:
				return THIRD;
			default:
				return null;
		}
	}
}
