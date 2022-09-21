import { useQuery } from '@tanstack/react-query';
import { MainContext } from 'context/provider';
import Data from 'models/data';
import User from 'models/user';
import { FC, ReactElement, useContext } from 'react';
import { Navigate } from 'react-router-dom';
import { RequestToServer } from 'services/Api';

const Public: FC<{ children: ReactElement }> = ({ children }) => {
  const { user, setUser } = useContext(MainContext);

  useQuery<Data<User>>(['user'], () => RequestToServer({ url: 'getuser' }), {
    onSuccess: () => {
      setUser(true);
    },
  });
  if (user) {
    return <Navigate to={'/'} />;
  }

  return <>{children}</>;
};

export default Public;
