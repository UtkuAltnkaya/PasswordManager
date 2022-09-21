import { useQuery } from '@tanstack/react-query';
import Loader from 'components/Loader/Loader';
import { MainContext } from 'context/provider';
import Data from 'models/data';
import User from 'models/user';
import { FC, useContext } from 'react';
import { Navigate } from 'react-router-dom';
import { RequestToServer } from 'services/Api';

const Private: FC<{ children: Function }> = ({ children }) => {
  const { setUser } = useContext(MainContext);
  const { data, isLoading, isError } = useQuery<Data<User>>(['user'], () => RequestToServer({ url: 'getuser' }), {
    onSuccess: () => {
      setUser(true);
    },
  });

  if (isLoading) {
    return <Loader />;
  }

  if (isError) {
    return <Navigate to={'/login'} />;
  }

  return <>{children(data)}</>;
};

export default Private;
