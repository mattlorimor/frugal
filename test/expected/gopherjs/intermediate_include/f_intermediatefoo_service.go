// Autogenerated by Frugal Compiler (3.11.1)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package intermediate_include

import (
	"fmt"

	"github.com/Workiva/frugal/lib/gopherjs/frugal"
	"github.com/Workiva/frugal/lib/gopherjs/thrift"
)

type FIntermediateFoo interface {
	IntermeidateFoo(ctx frugal.FContext) (err error)
}

type FIntermediateFooClient struct {
	c       frugal.FClient
	methods map[string]*frugal.Method
}

func NewFIntermediateFooClient(provider *frugal.FServiceProvider, middleware ...frugal.ServiceMiddleware) *FIntermediateFooClient {
	methods := make(map[string]*frugal.Method)
	client := &FIntermediateFooClient{
		c:       frugal.NewFStandardClient(provider),
		methods: methods,
	}
	middleware = append(middleware, provider.GetMiddleware()...)
	methods["intermeidateFoo"] = frugal.NewMethod(client, client.intermeidateFoo, "intermeidateFoo", middleware)
	return client
}

func (f *FIntermediateFooClient) Client_() frugal.FClient { return f.c }

func (f *FIntermediateFooClient) IntermeidateFoo(ctx frugal.FContext) (err error) {
	ret := f.methods["intermeidateFoo"].Invoke([]interface{}{ctx})
	if len(ret) != 1 {
		panic(fmt.Sprintf("Middleware returned %d arguments, expected 1", len(ret)))
	}
	if ret[0] != nil {
		err = ret[0].(error)
	}
	return err
}

func (f *FIntermediateFooClient) intermeidateFoo(ctx frugal.FContext) (err error) {
	args := IntermediateFooIntermeidateFooArgs{}
	result := IntermediateFooIntermeidateFooResult{}
	err = f.Client_().Call(ctx, "intermeidateFoo", &args, &result)
	if err != nil {
		return
	}
	return
}

type FIntermediateFooProcessor struct {
	*frugal.FBaseProcessor
}

func NewFIntermediateFooProcessor(handler FIntermediateFoo, middleware ...frugal.ServiceMiddleware) *FIntermediateFooProcessor {
	p := &FIntermediateFooProcessor{frugal.NewFBaseProcessor()}
	p.AddToProcessorMap("intermeidateFoo", &intermediatefooFIntermeidateFoo{frugal.NewFBaseProcessorFunction(p.GetWriteMutex(), frugal.NewMethod(handler, handler.IntermeidateFoo, "IntermeidateFoo", middleware))})
	return p
}

type intermediatefooFIntermeidateFoo struct {
	*frugal.FBaseProcessorFunction
}

func (p *intermediatefooFIntermeidateFoo) Process(ctx frugal.FContext, iprot, oprot *frugal.FProtocol) error {
	args := IntermediateFooIntermeidateFooArgs{}
	err := args.Read(iprot)
	iprot.ReadMessageEnd()
	if err != nil {
		return p.SendError(ctx, oprot, frugal.APPLICATION_EXCEPTION_PROTOCOL_ERROR, "intermeidateFoo", err.Error())
	}
	result := IntermediateFooIntermeidateFooResult{}
	ret := p.InvokeMethod([]interface{}{ctx})
	if len(ret) != 1 {
		panic(fmt.Sprintf("Middleware returned %d arguments, expected 1", len(ret)))
	}
	if ret[0] != nil {
		err = ret[0].(error)
	}
	if err != nil {
		if typedError, ok := err.(thrift.TApplicationException); ok {
			p.SendError(ctx, oprot, typedError.TypeId(), "intermeidateFoo", typedError.Error())
			return nil
		}
		return p.SendError(ctx, oprot, frugal.APPLICATION_EXCEPTION_INTERNAL_ERROR, "intermeidateFoo", "Internal error processing intermeidateFoo: "+err.Error())
	}
	return p.SendReply(ctx, oprot, "intermeidateFoo", &result)
}

type IntermediateFooIntermeidateFooArgs struct {
}

func NewIntermediateFooIntermeidateFooArgs() *IntermediateFooIntermeidateFooArgs {
	return &IntermediateFooIntermeidateFooArgs{}
}

func (p *IntermediateFooIntermeidateFooArgs) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		if err := iprot.Skip(fieldTypeId); err != nil {
			return err
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *IntermediateFooIntermeidateFooArgs) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("IntermeidateFoo_args"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *IntermediateFooIntermeidateFooArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("IntermediateFooIntermeidateFooArgs(%+v)", *p)
}

type IntermediateFooIntermeidateFooResult struct {
}

func NewIntermediateFooIntermeidateFooResult() *IntermediateFooIntermeidateFooResult {
	return &IntermediateFooIntermeidateFooResult{}
}

func (p *IntermediateFooIntermeidateFooResult) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		if err := iprot.Skip(fieldTypeId); err != nil {
			return err
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *IntermediateFooIntermeidateFooResult) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("IntermeidateFoo_result"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *IntermediateFooIntermeidateFooResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("IntermediateFooIntermeidateFooResult(%+v)", *p)
}
