import { HttpInterceptorFn } from '@angular/common/http';

export const credentialInterceptor: HttpInterceptorFn = (req, next) => {
  const r = req.clone({
    withCredentials: true,
  });

  return next(r);
};
