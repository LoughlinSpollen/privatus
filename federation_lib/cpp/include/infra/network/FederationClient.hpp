#ifndef MOTORCYCLE_H
#define MOTORCYCLE_H

// Observer pattern - callback function from C++ to Python
using FederationStateChangeObserver = std::function<bool(std::string, long)>;

namespace network {

class FederationClient {
public:
    FederationClient(FederationStateChangeObserver &observer);
    ~FederationClient();

private:
    FederationStateChangeObserver &m_Observer;
};

}

#endif
