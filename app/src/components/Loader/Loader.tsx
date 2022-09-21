import style from './Loader.module.scss';

type Props = {
  width?: string;
  height?: string;
  border?: string;
  borderTop?: string;
};

const Loader = ({ width = '50px', height = '50px', border = '10px', borderTop = '10px' }: Props) => {
  return (
    <div className={style.main}>
      <div
        className={style.loadingSpinner}
        style={{ width, height, border: `${border} solid #f3f3f3`, borderTop: `${borderTop} solid #383636` }}
      />
    </div>
  );
};

export default Loader;
