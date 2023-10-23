#include "Federation.h"

namespace privatus {

Federation::Federation(const std::string &modelName) : m_ModelName(modelName) {

}

Federation::~Federation() {

}

void Federation::Register() {

}

FederationConfig& Federation::GetFederationConfig() const {
    return *m_FederationConfig;
}

}
