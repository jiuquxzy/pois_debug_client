package client

import (
	"context"
	"log"
	"pois_debug_client/config"
	"pois_debug_client/pois_rpc"
	"time"

	"github.com/CESSProject/cess_pois/acc"
	"github.com/CESSProject/cess_pois/pois"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func SendCommits(minerId []byte, key acc.RsaKey, front, rear int64, acc []byte, commits pois.Commits) ([][]int64, error) {
	conn, err := grpc.Dial(config.GetConfig().TeeAddr,
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	cli := pois_rpc.NewPoisApiClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	tsCommits := pois_rpc.Commits{
		FileIndexs: commits.FileIndexs,
		Roots:      commits.Roots,
	}
	response, err := cli.RequestMinerCommitGenChall(
		ctx, &pois_rpc.RequestMinerCommitGenChall{
			MinerId: minerId,
			Commit:  &tsCommits,
			PoisInfo: &pois_rpc.MinerPoisInfo{
				Acc:   acc,
				Front: front,
				Rear:  rear,
				KeyN:  key.N.Bytes(),
				KeyG:  key.G.Bytes(),
			},
		},
	)
	if err != nil {
		return nil, err
	}

	chals := make([][]int64, len(response.Rows))
	for i := 0; i < len(chals); i++ {
		chals[i] = response.Rows[i].Values
	}
	return chals, nil
}

func SendCommitProofs(minerId []byte, commitProofs [][]pois.CommitProof, accProofs *pois.AccProof) (bool, error) {
	maxMsgSize := 1024 * 1024 * 100
	conn, err := grpc.Dial(
		config.GetConfig().TeeAddr,
		grpc.WithTransportCredentials(
			insecure.NewCredentials(),
		),
		grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(maxMsgSize),
			grpc.MaxCallSendMsgSize(maxMsgSize),
		),
	)
	if err != nil {
		return false, err
	}
	defer conn.Close()
	cli := pois_rpc.NewPoisApiClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	commitGroup := &pois_rpc.CommitProofGroup{
		CommitProofGroupInner: make([]*pois_rpc.CommitProofGroupInner, len(commitProofs)),
	}
	for i := 0; i < len(commitProofs); i++ {
		commitGroup.CommitProofGroupInner[i] = &pois_rpc.CommitProofGroupInner{
			CommitProof: make([]*pois_rpc.CommitProof, len(commitProofs[i])),
		}
		for j := 0; j < len(commitProofs[i]); j++ {
			commitGroup.CommitProofGroupInner[i].CommitProof[j] = &pois_rpc.CommitProof{
				Node: &pois_rpc.MhtProof{
					Index: int32(commitProofs[i][j].Node.Index),
					Label: commitProofs[i][j].Node.Label,
					Paths: commitProofs[i][j].Node.Paths,
					Locs:  commitProofs[i][j].Node.Locs,
				},
				Parents: make([]*pois_rpc.MhtProof, len(commitProofs[i][j].Parents)),
				Elders:  make([]*pois_rpc.MhtProof, len(commitProofs[i][j].Elders)),
			}
			for k := 0; k < len(commitProofs[i][j].Parents); k++ {
				commitGroup.CommitProofGroupInner[i].CommitProof[j].Parents[k] = &pois_rpc.MhtProof{
					Index: int32(commitProofs[i][j].Parents[k].Index),
					Label: commitProofs[i][j].Parents[k].Label,
					Paths: commitProofs[i][j].Parents[k].Paths,
					Locs:  commitProofs[i][j].Parents[k].Locs,
				}
			}
			for k := 0; k < len(commitProofs[i][j].Elders); k++ {
				commitGroup.CommitProofGroupInner[i].CommitProof[j].Elders[k] = &pois_rpc.MhtProof{
					Index: int32(commitProofs[i][j].Elders[k].Index),
					Label: commitProofs[i][j].Elders[k].Label,
					Paths: commitProofs[i][j].Elders[k].Paths,
					Locs:  commitProofs[i][j].Elders[k].Locs,
				}
			}
		}
	}

	accProof := &pois_rpc.AccProof{
		Indexs:  accProofs.Indexs,
		Labels:  accProofs.Labels,
		AccPath: accProofs.AccPath,
		WitChains: &pois_rpc.AccWitnessNode{
			Elem: accProofs.WitChains.Elem,
			Wit:  accProofs.WitChains.Wit,
			Acc: &pois_rpc.AccWitnessNode{
				Elem: accProofs.WitChains.Acc.Elem,
				Wit:  accProofs.WitChains.Acc.Wit,
				Acc: &pois_rpc.AccWitnessNode{
					Elem: accProofs.WitChains.Acc.Acc.Elem,
					Wit:  accProofs.WitChains.Acc.Acc.Wit,
				},
			},
		},
	}

	response, err := cli.RequestVerifyCommitProof(ctx,
		&pois_rpc.RequestVerifyCommitAndAccProof{
			MinerId:          minerId,
			CommitProofGroup: commitGroup,
			AccProof:         accProof,
		},
	)
	if err != nil {
		return false, err
	}
	log.Println("response", response.PoisStatus)
	return true, nil
}

func SendSpaceProof(minerId []byte, chal []int64, key acc.RsaKey, front, rear int64, acc []byte, spaceProof *pois.SpaceProof) (bool, error) {

	maxMsgSize := 1024 * 1024 * 100
	conn, err := grpc.Dial(
		config.GetConfig().TeeAddr,
		grpc.WithTransportCredentials(
			insecure.NewCredentials(),
		),
		grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(maxMsgSize),
			grpc.MaxCallSendMsgSize(maxMsgSize),
		),
	)
	if err != nil {
		return false, err
	}
	defer conn.Close()
	cli := pois_rpc.NewPoisApiClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	mhtProofGroup := make([]*pois_rpc.MhtProofGroup, len(spaceProof.Proofs))

	for i := 0; i < len(spaceProof.Proofs); i++ {
		mhtProofGroup[i] = &pois_rpc.MhtProofGroup{}
		mhtProofGroup[i].Proofs = make([]*pois_rpc.MhtProof, len(spaceProof.Proofs[i]))
		for j := 0; j < len(spaceProof.Proofs[i]); j++ {
			mhtProofGroup[i].Proofs[j] = &pois_rpc.MhtProof{
				Index: int32(spaceProof.Proofs[i][j].Index),
				Label: spaceProof.Proofs[i][j].Label,
				Paths: spaceProof.Proofs[i][j].Paths,
				Locs:  spaceProof.Proofs[i][j].Locs,
			}
		}
	}

	var witChains = make([]*pois_rpc.AccWitnessNode, len(spaceProof.WitChains))

	for i := 0; i < len(spaceProof.WitChains); i++ {
		witChains[i] = &pois_rpc.AccWitnessNode{
			Elem: spaceProof.WitChains[i].Elem,
			Wit:  spaceProof.WitChains[i].Wit,
			Acc: &pois_rpc.AccWitnessNode{
				Elem: spaceProof.WitChains[i].Acc.Elem,
				Wit:  spaceProof.WitChains[i].Acc.Wit,
				Acc: &pois_rpc.AccWitnessNode{
					Elem: spaceProof.WitChains[i].Acc.Acc.Elem,
					Wit:  spaceProof.WitChains[i].Acc.Acc.Wit,
					Acc: &pois_rpc.AccWitnessNode{
						Elem: spaceProof.WitChains[i].Acc.Acc.Acc.Elem,
						Wit:  spaceProof.WitChains[i].Acc.Acc.Acc.Wit,
					},
				},
			},
		}
	}

	proof := &pois_rpc.SpaceProof{
		Left:      spaceProof.Left,
		Right:     spaceProof.Right,
		Roots:     spaceProof.Roots,
		Proofs:    mhtProofGroup,
		WitChains: witChains,
	}
	poisInfo := &pois_rpc.MinerPoisInfo{
		Acc:   acc,
		Front: front,
		Rear:  rear,
		KeyN:  key.N.Bytes(),
		KeyG:  key.G.Bytes(),
	}

	resp, err := cli.RequestSpaceProofVerifySingleBlock(ctx,
		&pois_rpc.RequestSpaceProofVerify{
			SpaceChals: chal,
			MinerId:    minerId,
			PoisInfo:   poisInfo,
			Proof:      proof,
		})

	if err != nil {
		return false, err
	}
	log.Println("response", resp.Signature)
	return true, nil
}

func SendDeletionProof(minerId []byte, key acc.RsaKey, front, rear int64, acc []byte, delProof *pois.DeletionProof) (bool, error) {

	maxMsgSize := 1024 * 1024 * 100
	conn, err := grpc.Dial(
		config.GetConfig().TeeAddr,
		grpc.WithTransportCredentials(
			insecure.NewCredentials(),
		),
		grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(maxMsgSize),
			grpc.MaxCallSendMsgSize(maxMsgSize),
		),
	)
	if err != nil {
		return false, err
	}
	defer conn.Close()
	cli := pois_rpc.NewPoisApiClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	witChain := &pois_rpc.AccWitnessNode{
		Elem: delProof.WitChain.Elem,
		Wit:  delProof.WitChain.Wit,
		Acc: &pois_rpc.AccWitnessNode{
			Elem: delProof.WitChain.Acc.Elem,
			Wit:  delProof.WitChain.Acc.Wit,
			Acc: &pois_rpc.AccWitnessNode{
				Elem: delProof.WitChain.Acc.Acc.Elem,
				Wit:  delProof.WitChain.Acc.Acc.Wit,
			},
		},
	}
	requestVerifyDeletionProof := &pois_rpc.RequestVerifyDeletionProof{
		Roots:    delProof.Roots,
		WitChain: witChain,
		AccPath:  delProof.AccPath,
		MinerId:  minerId,
		PoisInfo: &pois_rpc.MinerPoisInfo{
			Acc:   acc,
			Front: front,
			Rear:  rear,
			KeyN:  key.N.Bytes(),
			KeyG:  key.G.Bytes(),
		},
	}

	resp, err := cli.RequestVerifyDeletionProof(ctx, requestVerifyDeletionProof)

	if err != nil {
		return false, err
	}

	log.Println("response", resp.PoisStatus.Front)

	return true, nil
}
