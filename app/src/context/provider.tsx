import { useContext, useState, FC, ReactElement, createContext } from 'react';

interface ContextInterface {
  user: boolean;
  setUser?: (item: boolean) => void;
}

const defaultState = {
  user: false,
};

const MainContext = createContext<ContextInterface>(defaultState);

const MainProvider: FC<{ children: ReactElement }> = ({ children }) => {
  const [user, setUser] = useState<boolean>(false);

  const data = {
    user,
    setUser,
  };

  return <MainContext.Provider value={data}>{children}</MainContext.Provider>;
};

export { useContext, MainProvider, MainContext };
