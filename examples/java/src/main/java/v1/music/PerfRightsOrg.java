/**
 * Autogenerated by Frugal Compiler (2.15.0)
 * DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING
 *
 * @generated
 */
package v1.music;

import java.util.Map;
import java.util.HashMap;
import org.apache.thrift.TEnum;

public enum PerfRightsOrg implements org.apache.thrift.TEnum {
	ASCAP(1),
	BMI(2),
	SESAC(3),
	Other(4);

	private final int value;

	private PerfRightsOrg(int value) {
		this.value = value;
	}

	public int getValue() {
		return value;
	}

	public static PerfRightsOrg findByValue(int value) {
		switch (value) {
			case 1:
				return ASCAP;
			case 2:
				return BMI;
			case 3:
				return SESAC;
			case 4:
				return Other;
			default:
				return null;
		}
	}
}
