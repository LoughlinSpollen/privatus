#include "../../../include/infra/network/FederationClient.hpp"

#include <iostream>

namespace network {

FederationClient::FederationClient(FederationStateChangeObserver &observer) :
    m_Observer(observer) {
}

FederationClient::~FederationClient() {

}

}
