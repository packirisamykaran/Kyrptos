import { error } from "console";
import { Whale } from "../../models/whaleModels";
import { defiAddress } from "./constants";

function getWhalesByProtocol(protocol: string): Whale[] {
  let whales: Whale[] = [];
  try {
    let protocolAddress = defiAddress.get(protocol);

    if (protocolAddress !== undefined) {
      throw new Error("Protocol address not found");
    }

    let params: Record<string, any> = {
      tokenAddress: protocolAddress,
      limit: 5,
      offset: 0,
    };
  } catch (error) {
    throw error;
  }

  return whales;
}
